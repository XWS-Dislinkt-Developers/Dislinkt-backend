package config

type Config struct {
	Port                   string
	AuthenticationHost     string
	AuthenticationPort     string
	UserPostHost           string
	UserPostPort           string
	UserConnectionHost     string
	UserConnectionPort     string
	UserHost               string
	UserPort               string
	HTTPSServerKey         string
	HTTPSServerCertificate string
}

func NewConfig() *Config {
	return &Config{
		Port: "8000",

		AuthenticationHost: "localhost",
		AuthenticationPort: "8001",

		UserPostHost: "localhost",
		UserPostPort: "8002",

		UserHost: "localhost",
		UserPort: "8003",

		UserConnectionHost: "localhost",
		UserConnectionPort: "8004",

		HTTPSServerKey:         ".cert/server.key",
		HTTPSServerCertificate: ".cert/server.crt",

		//DOCKER
		//Port:                   os.Getenv("GATEWAY_PORT"),
		//AuthenticationHost:     os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		//AuthenticationPort:     os.Getenv("AUTHENTICATION_SERVICE_PORT"),
		//UserHost:               os.Getenv("USER_SERVICE_HOST"),
		//UserPort:               os.Getenv("USER_SERVICE_PORT"),
		//HTTPSServerKey:         os.Getenv("HTTPS_SERVER_KEY"),
		//HTTPSServerCertificate: os.Getenv("HTTPS_SERVER_CERTIFICATE"),
	}
}
