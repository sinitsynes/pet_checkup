package app

import "github.com/jackc/pgx/v5/pgxpool"

type Application struct {
	DbPool *pgxpool.Pool
}
