package main

import (
	"context"
	"log/slog"
	"os"

	"RNTreaty/cmd/api"
	"RNTreaty/model"

	databaseConfig "RNTreaty/db"

	"github.com/joho/godotenv"
)

func main() {
	slog.Info("Server Starting")
	if err := godotenv.Load(".env"); err != nil {
		slog.Error(err.Error())
	}

	// Port initialization
	port := os.Getenv("PORT")
	if port == "" {
		slog.Error("PORT NOT FOUND")
	}

	// Database initialization
	dbConnection := os.Getenv("DATABASE_URL")
	if dbConnection == "" {
		slog.Error("DATABASE URL NOT FOUND")
	}

	// Host initialization
	host := os.Getenv("HOST")
	if host == "" {
		slog.Error("HOST not found")
	}

	cfg := &model.Config{
		Host:          host,
		Port:          port,
		ConnectionURL: dbConnection,
	}

	// Context initialization
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Database initialization
	db := databaseConfig.ConnectDB(ctx, cfg)
	defer db.Close()

	// Server Loop
	app := api.NewTreatyAPI(cfg, db)
	if runErr := app.Run(); runErr != nil {
		slog.Error(runErr.Error())
	}
}
