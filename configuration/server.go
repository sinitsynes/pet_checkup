package configuration

type Server struct {
	Stand string
	Host  string
	Port  string
}

func newServer() Server {
	return Server{
		Stand: getEnvAsString("STAND", "local"),
		Host:  getEnvAsString("HOST", "0.0.0.0"),
		Port:  getEnvAsString("PORT", "8000"),
	}
}
