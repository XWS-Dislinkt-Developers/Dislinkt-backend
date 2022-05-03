package config

import "os"

type Config struct {
	Port               string
	AuthenticationHost string
	AuthenticationPort string
}

func NewConfig() *Config {
	return &Config{
		Port:               os.Getenv("GATEWAY_PORT"),                // "8000",      //os.Getenv("GATEWAY_PORT"),
		AuthenticationHost: os.Getenv("AUTHENTICATION_SERVICE_HOST"), //"localhost", //os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		AuthenticationPort: os.Getenv("AUTHENTICATION_SERVICE_PORT"), //"8001",      //os.Getenv("AUTHENTICATION_SERVICE_PORT"),
	}
}
