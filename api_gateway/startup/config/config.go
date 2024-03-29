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
	JobServiceHost         string
	JobServicePort         string
	MessageHost            string
	MessagePort            string
	HTTPSServerKey         string
	HTTPSServerCertificate string
}

func NewConfig() *Config {
	return &Config{

		//Port: "8000",
		//
		//AuthenticationHost: "localhost",
		//AuthenticationPort: "8001",
		//
		//UserPostHost: "localhost",
		//UserPostPort: "8002",
		//
		//UserHost: "localhost",
		//UserPort: "8003",
		//
		//UserConnectionHost: "localhost",
		//UserConnectionPort: "8004",
		//
		//JobServiceHost: "localhost",
		//JobServicePort: "8005",
		//
		//MessageHost: "localhost",
		//MessagePort: "8008",
		//
		//HTTPSServerKey:         ".cert/server.key",
		//HTTPSServerCertificate: ".cert/server.crt",

		//DOCKER

		Port:                   os.Getenv("GATEWAY_PORT"),
		AuthenticationHost:     os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		AuthenticationPort:     os.Getenv("AUTHENTICATION_SERVICE_PORT"),
		UserHost:               os.Getenv("USER_SERVICE_HOST"),
		UserPort:               os.Getenv("USER_SERVICE_PORT"),
		UserConnectionHost:     os.Getenv("USER_CONNECTION_SERVICE_HOST"),
		UserConnectionPort:     os.Getenv("USER_CONNECTION_SERVICE_PORT"),
		HTTPSServerKey:         os.Getenv("HTTPS_SERVER_KEY"),
		HTTPSServerCertificate: os.Getenv("HTTPS_SERVER_CERTIFICATE"),
		UserPostHost:           os.Getenv("USER_POST_SERVICE_HOST"),
		UserPostPort:           os.Getenv("USER_POST_SERVICE_PORT"),

		JobServiceHost: os.Getenv("JOB_SERVICE_HOST"),
		JobServicePort: os.Getenv("JOB_SERVICE_PORT"),

		MessageHost: os.Getenv("MESSAGE_SERVICE_HOST"),
		MessagePort: os.Getenv("MESSAGE_SERVICE_PORT"),
	}
}
