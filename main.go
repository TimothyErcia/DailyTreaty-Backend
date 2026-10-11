package main

import (
	"context"
	"log/slog"
	"os"

	"RNTreaty/cmd/api"
	"RNTreaty/handlers"
	"RNTreaty/model"

	databaseConfig "RNTreaty/db"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	slog.Info("Server Starting")
	if err := godotenv.Load(".env"); err != nil {
		slog.Error(err.Error())
	}

	// Config initialization
	port := os.Getenv("PORT")
	if port == "" {
		slog.Error("PORT NOT FOUND")
	}

	dbConnection := os.Getenv("DATABASE_URL")
	if dbConnection == "" {
		slog.Error("DATABASE URL NOT FOUND")
	}

	cfg := &model.Config{
		Host:          "localhost",
		Port:          port,
		ConnectionURL: dbConnection,
	}

	// Context initialization
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Database initialization
	db := databaseConfig.ConnectDB(ctx, cfg)
	defer db.Close()

	// Repo Layer
	// treatyRepo := NewTreatyRepo(db)

	// Usecase Layer
	// usecase := NewTreatyUseCase(treatyRepo)

	// Handlers
	route := chi.NewRouter()
	route = handlers.NewHealthzHandler(route)

	// Server Loop
	app := api.NewTreatyAPI(cfg, db, route)
	if runErr := app.Run(); runErr != nil {
		slog.Error(runErr.Error())
	}
}
