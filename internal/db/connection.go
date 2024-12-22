package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// DatabaseConfig is configuration wrapper for our database setting
type DatabaseConfig struct {
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	Name     string `env:"NAME"`
}

func ConnectionString() (string, error) {
	loadEnvErr := godotenv.Load()
	if loadEnvErr != nil {
		return "", errors.New("can't load .env")
	}
	db := DatabaseConfig{}
	parseErr := env.Parse(&db, env.Options{Prefix: "DB_"})
	if parseErr != nil {
		return "", parseErr
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		db.User, db.Password, db.Host, db.Port, db.Name), nil
}

// NewDBPool will create pool connection to database
func NewDBPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
