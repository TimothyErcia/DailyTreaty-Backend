package db

import (
	"context"
	"log"

	"RNTreaty/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(ctx context.Context, cfg *model.Config) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, cfg.ConnectionURL)
	if err != nil {
		log.Fatal(err)
	}

	return pool
}
