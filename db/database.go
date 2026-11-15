package db

import (
	"context"
	"fmt"
	"log/slog"

	"RNTreaty/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(ctx context.Context, cfg *model.Config) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, cfg.ConnectionURL)
	if err != nil {
		slog.Error(err.Error())
	}

	if pingErr := pool.Ping(ctx); pingErr != nil {
		slog.Error(pingErr.Error())
	}

	slog.Info(fmt.Sprintf("Listening to %s at port %s\n", cfg.Host, cfg.Port))
	return pool
}
