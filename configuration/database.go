package configuration

import "fmt"

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func (db DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		db.Username, db.Password, db.Host, db.Port, db.Name)
}

func newDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Username: getEnvAsString("POSTGRES_USER", ""),
		Password: getEnvAsString("POSTGRES_PASSWORD", ""),
		Host:     getEnvAsString("POSTGRES_HOST", ""),
		Port:     getEnvAsString("POSTGRES_PORT", ""),
		Name:     getEnvAsString("POSTGRES_DB", ""),
	}
}
