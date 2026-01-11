package api

import (
	"reflect"
	"strings"

	"github.com/ankitsalunkhe/chat-app-backend/internal/api/gen"
	"github.com/ankitsalunkhe/chat-app-backend/internal/auth"
	"github.com/ankitsalunkhe/chat-app-backend/internal/storage"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var _ gen.ServerInterface = (*Handlers)(nil)

type Handlers struct {
	db   storage.Database
	cv   CustomValidator
	auth auth.Auth
}

func New(auth auth.Auth, db storage.Database) Handlers {
	validator := validator.New()
	cv := CustomValidator{
		validator: validator,
	}

	// register function to get tag name from json tags.
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return Handlers{db: db, cv: cv, auth: auth}
}

func (h Handlers) GetHealth(ctx echo.Context) error {
	return nil
}
