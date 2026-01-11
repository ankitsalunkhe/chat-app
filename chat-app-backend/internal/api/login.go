package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ankitsalunkhe/chat-app-backend/internal/api/gen"
	"github.com/ankitsalunkhe/chat-app-backend/internal/storage"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h Handlers) PostLogin(c echo.Context) error {
	requestBody := gen.LoginRequest{}

	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			gen.ErrorResponse{
				Error:            gen.InvalidRequest,
				ErrorDescription: "binding request: " + err.Error(),
			},
		)
	}
	if err := h.cv.Validate(&requestBody); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			gen.ErrorResponse{
				Error:            gen.InvalidRequest,
				ErrorDescription: "invalid request: " + err.Error(),
			},
		)
	}

	user, err := h.db.GetUser(requestBody.Username)
	if err != nil {
		if errors.Is(err, storage.ErrNoSuchUsername) {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				fmt.Errorf("no such username exists"),
			)
		}
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Errorf("querying database for password: %w", err),
		)
	}

	match := checkPasswordHash(requestBody.Password, user.Password)
	if !match {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			fmt.Errorf("wrong password"),
		)
	}

	if err := h.auth.SetSession(c, user.Id); err != nil {
		return fmt.Errorf("error setting session :%w", err)
	}

	return nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (h Handlers) PostLogout(e echo.Context) error {
	fmt.Println("cookies", e.Cookies())
	return h.auth.RemoveSession(e)
}
