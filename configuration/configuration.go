package configuration

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	DB     DatabaseConfig
	Server Server
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	return &Config{
		DB:     newDatabaseConfig(),
		Server: newServer(),
	}
}
