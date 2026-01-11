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

func (h Handlers) PostRegister(c echo.Context) error {
	requestBody := gen.RegisterRequest{}
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			gen.ErrorResponse{
				Error:            gen.InvalidRequest,
				ErrorDescription: "binding request: " + err.Error(),
			},
		)
	}

	if err := h.cv.Validate(requestBody); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			gen.ErrorResponse{
				Error:            gen.InvalidRequest,
				ErrorDescription: "invalid request: " + err.Error(),
			},
		)
	}

	hashedPassword, err := hashPassword(requestBody.Password)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Errorf("hashing password from request: %w", err),
		)
	}

	if err := h.db.CreateCredentials(requestBody.Username, hashedPassword); err != nil {
		if errors.Is(err, storage.ErrDuplicateUsername) {
			return echo.NewHTTPError(
				http.StatusConflict,
				fmt.Errorf("username already exists"),
			)
		}
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, gen.RegisterResponse{
		Success: true,
	})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
