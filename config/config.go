package config

import (
	"errors"
	"os"
)

// Config holds application configuration loaded from environment variables.
type Config struct {
	DatabaseURL string
	JwtSecret   string
	ServerPort  string
	SMTPHost    string
	SMTPPort    string
	SMTPUser    string
	SMTPPass    string
	SMTPFrom    string
}

// Load reads configuration from environment variables. Minimal validation is done.
func Load() (*Config, error) {
	cfg := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
		ServerPort:  os.Getenv("SERVER_PORT"),
		SMTPHost:    os.Getenv("SMTP_HOST"),
		SMTPPort:    os.Getenv("SMTP_PORT"),
		SMTPUser:    os.Getenv("SMTP_USER"),
		SMTPPass:    os.Getenv("SMTP_PASS"),
		SMTPFrom:    os.Getenv("SMTP_FROM"),
	}

	if cfg.DatabaseURL == "" || cfg.JwtSecret == "" {
		return nil, errors.New("DATABASE_URL and JWT_SECRET are required")
	}
	if cfg.ServerPort == "" {
		cfg.ServerPort = ":8080"
	}

	return cfg, nil
}
