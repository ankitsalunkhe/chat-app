package api

import (
	"fmt"
	"net/http"

	"github.com/ankitsalunkhe/chat-app-backend/internal/api/gen"
	"github.com/labstack/echo/v4"
)

func (h Handlers) PostVerify(e echo.Context) error {
	userId, err := h.auth.VerifySession(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, gen.UnauthorizedResponse{
			Error: fmt.Sprintf("invalid session: %v", err),
		})
	}

	return e.JSON(200, gen.VerifyAuthResponse{
		UserId: int(userId),
	})
}
