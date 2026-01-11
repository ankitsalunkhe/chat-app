package auth

import (
	"errors"
	"log"

	"github.com/boj/redistore"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	rdstore *redistore.RediStore
}

const (
	userIdSession    = "userID"
	sessionCookieKey = "chatapp-session"
)

var (
	errSessionNotFound      = errors.New("session not found")
	errSessionUserNotFound  = errors.New("user id not found in session")
	errSessionUserNotString = errors.New("user id not string")
)

const (
	// cookieMaxAge is the maximum age of the session cookie.
	cookieMaxAge = 60 * 60 // 1 hour
	redisMaxAge  = 60 * 60 //  1 hour
)

func New(sessionKey string) Auth {

	rdstore, err := redistore.NewRediStore(10, "tcp", ":63798", "", "", []byte(sessionKey))
	if err != nil {
		panic(err)
	}

	rdstore.Options.MaxAge = cookieMaxAge
	rdstore.Options.HttpOnly = true
	rdstore.DefaultMaxAge = cookieMaxAge

	return Auth{
		rdstore: rdstore,
	}
}

func (a *Auth) SetSession(e echo.Context, userID int32) error {
	session, err := a.rdstore.Get(e.Request(), sessionCookieKey)
	if err != nil {
		log.Println(err.Error())
	}

	session.Values[userIdSession] = userID

	if err = session.Save(e.Request(), e.Response()); err != nil {
		log.Fatalf("Error saving session: %v", err)
	}

	return nil
}

func (a *Auth) VerifySession(e echo.Context) (int32, error) {
	session, err := a.rdstore.Get(e.Request(), sessionCookieKey)
	if err != nil {
		return 0, errSessionNotFound
	}

	userId, ok := session.Values[userIdSession]
	if !ok {
		return 0, errSessionUserNotFound
	}

	userIdString, ok := userId.(int32)
	if !ok {
		return 0, errSessionUserNotString
	}

	return userIdString, nil
}

func (a *Auth) RemoveSession(e echo.Context) error {
	session, err := a.rdstore.Get(e.Request(), sessionCookieKey)
	if err != nil {
		return errSessionNotFound
	}

	session.Options.MaxAge = -1

	if err = session.Save(e.Request(), e.Response()); err != nil {
		log.Fatalf("Error saving session: %v", err)
	}

	return nil
}
