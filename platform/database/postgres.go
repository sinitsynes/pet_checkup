package database

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		slog.Error("Unable to parse connection string",
			"value", connString,
			"err", err,
		)
	}
	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		slog.Error("Unable to create connection pool",
			"value", connString,
			"err", err,
		)
	}
	err = dbPool.Ping(context.Background())
	if err != nil {
		slog.Error("Unable to connect to database",
			"value", connString,
			"err", err,
		)
	}
	return dbPool, err
}
