package pg_connection

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPG(ctx context.Context) (*pgxpool.Pool, error) {
	connString := os.Getenv("DB_DSN")
	db, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("cannot create connection pool %q", err)
	}
	return db, err
}
