package main

import (
	"github.com/ankitsalunkhe/chat-app-backend/internal/api"
	"github.com/ankitsalunkhe/chat-app-backend/internal/api/gen"
	"github.com/ankitsalunkhe/chat-app-backend/internal/auth"
	"github.com/ankitsalunkhe/chat-app-backend/internal/config"
	"github.com/ankitsalunkhe/chat-app-backend/internal/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	auth := auth.New(cfg.SessionKey)

	db, err := storage.New(cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}

	handlers := api.New(auth, db)

	gen.RegisterHandlers(e, handlers)

	e.Logger.Fatal(e.Start(":1323"))
}
