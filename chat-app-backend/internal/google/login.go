package google

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ankitsalunkhe/chat-app-backend/internal/auth"
	"github.com/ankitsalunkhe/chat-app-backend/internal/storage"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	googleAPI "golang.org/x/oauth2/google"
)

const userInfoUrl = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type Google struct {
	config *oauth2.Config
	auth   auth.Auth
	db     storage.Database
}

func New(auth auth.Auth, db storage.Database) *Google {
	return &Google{
		config: &oauth2.Config{
			RedirectURL:  "http://localhost:1323/callback",
			ClientID:     "57789757605-gtdtc71c4nputu8g6rqm7tu7ut287s22.apps.googleusercontent.com",
			ClientSecret: "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",
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
			},
			Endpoint: googleAPI.Endpoint,
		},
		auth: auth,
		db:   db,
	}
}

func (google *Google) LoginHandler(c echo.Context) error {
	w := c.Response().Writer
	oauthState := generateStateOauthCookie(w)
	u := google.config.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return c.Redirect(http.StatusTemporaryRedirect, u)
}

func (google *Google) Callback(c echo.Context) error {
	r := c.Request()
	googleUser, err := google.getUserData(r.FormValue("code"))
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Errorf("getting google user data :%w", err),
		)
	}

	if !googleUser.VerifiedEmail {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Errorf("error signing in"),
		)

	}

	user, err := google.db.GetUser(googleUser.Email)
	if err == nil {
		if err := google.auth.SetSession(c, user.Id); err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError, fmt.Errorf("error setting session :%w", err),
			)
		}
		return nil
	}

	if errors.Is(err, storage.ErrNoSuchUsername) {
		if err := google.db.CreateGoogleCredential(googleUser.Email); err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				fmt.Errorf("saving google credential :%w", err),
			)
		}
	}

	if err := google.auth.SetSession(c, user.Id); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Errorf("setting session :%w", err),
		)
	}

	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:4000/home")
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return state
}

type googleUser struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

func (google *Google) getUserData(code string) (googleUser, error) {
	// Use code to get token and get user info from Google.
	token, err := google.config.Exchange(context.Background(), code)
	if err != nil {
		return googleUser{}, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	response, err := http.Get(userInfoUrl + token.AccessToken)
	if err != nil {
		return googleUser{}, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return googleUser{}, fmt.Errorf("failed read response: %s", err.Error())
	}

	var u googleUser
	if err := json.Unmarshal(contents, &u); err != nil {
		return googleUser{}, fmt.Errorf("marshalling user infor: %w", err)
	}

	return u, nil
}
