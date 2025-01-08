package app

import (
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

type DatabaseConfig struct {
	User             string
	Password         string
	Host             string
	Port             string
	Name             string
	ConnectionString string
}

type Configuration struct {
	DB *DatabaseConfig
}

type Application struct {
	Config *Configuration
	DbPool *pgxpool.Pool
}

func configurationSetup() *Configuration {
	cfg := &Configuration{
		DB: &DatabaseConfig{
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Name:     os.Getenv("POSTRES_DB"),
		},
	}
	cfg.DB.ConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	return cfg
}

func ApplicationSetup() *Application {
	return &Application{
		Config: configurationSetup(),
	}
}
