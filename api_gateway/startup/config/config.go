package config

type Config struct {
	Port               string
	AuthenticationHost string
	AuthenticationPort string
	UserPostHost       string
	UserPostPort       string
}

func NewConfig() *Config {
	return &Config{
		Port: "8000", //os.Getenv("GATEWAY_PORT"),

		AuthenticationHost: "localhost", //os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		AuthenticationPort: "8001",      //os.Getenv("AUTHENTICATION_SERVICE_PORT"),

		UserPostHost: "localhost",
		UserPostPort: "8002",
	}
}
