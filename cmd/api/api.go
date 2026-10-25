package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"RNTreaty/handlers"
	"RNTreaty/model"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Config   *model.Config
	Database *pgxpool.Pool
	Server   *http.Server
}

func NewTreatyAPI(cfg *model.Config, db *pgxpool.Pool) *App {
	// Repo Layer
	// treatyRepo := NewTreatyRepo(db)

	// Usecase Layer
	// usecase := NewTreatyUseCase(treatyRepo)

	// Route
	route := chi.NewRouter()
	route.Route("/v1", func(r chi.Router) {
		r.Mount("/healthz", handlers.NewHealthzHandler())
	})

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: route,
	}

	app := &App{
		Config:   cfg,
		Database: db,
		Server:   server,
	}

	return app
}

func (app *App) Run() error {
	fmt.Printf("Listening to %s at port %s\n", app.Config.Host, app.Config.Port)
	if err := app.Server.ListenAndServe(); err != nil {
		return err
	}

	// Server graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if srvErr := app.Server.Shutdown(ctx); srvErr != nil {
		return srvErr
	}

	return nil
}
