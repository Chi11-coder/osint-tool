//go:build dev

package main

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		slog.Info("not founded .env file, please check .env.example and configure your .env file")
	}
}
