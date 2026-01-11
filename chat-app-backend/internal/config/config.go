package config

import (
	"fmt"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
	SessionKey  string `envconfig:"SESSION_KEY"`
}

func New() (Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Debug("Error loading .env file")
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("unable to parse config variables: %w", err)
	}

	return cfg, nil
}
