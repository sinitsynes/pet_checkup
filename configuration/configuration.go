package configuration

type Config struct {
	DB     DatabaseConfig
	Server Server
}

func LoadConfig() *Config {
	return &Config{
		DB:     newDatabaseConfig(),
		Server: newServer(),
	}
}
