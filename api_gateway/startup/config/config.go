package config

import "os"

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
		Port: "8000", //os.Getenv("GATEWAY_PORT"),

		AuthenticationHost: "localhost", //os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		AuthenticationPort: "8001",      //os.Getenv("AUTHENTICATION_SERVICE_PORT"),

		UserPostHost: "localhost",
		UserPostPort: "8002",

		UserConnectionHost: "localhost",
		UserConnectionPort: "8004",

		UserHost: "localhost",
		UserPort: "8005",

		HTTPSServerKey:         os.Getenv("HTTPS_SERVER_KEY"),         //cert/"server.key",
		HTTPSServerCertificate: os.Getenv("HTTPS_SERVER_CERTIFICATE"), //"cert/server.crt",
	}
}
