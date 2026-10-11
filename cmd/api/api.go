package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"RNTreaty/model"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Config   *model.Config
	Database *pgxpool.Pool
	Server   *http.Server
}

func NewTreatyAPI(cfg *model.Config, db *pgxpool.Pool, handlers *chi.Mux) *App {
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handlers,
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
