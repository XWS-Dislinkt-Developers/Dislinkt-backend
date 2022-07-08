package config

import "os"

type Config struct {
	Port           string
	UserPostDBHost string
	UserPostDBPort string
	NatsHost       string
	NatsPort       string
	NatsUser       string
	NatsPass       string
}

func NewConfig() *Config {
	return &Config{
		//Port:           "8002",
		//UserPostDBHost: "localhost",
		//UserPostDBPort: "27017",
		//NatsHost:       "nats",
		//NatsPort:       "4222",
		//NatsUser:       "ruser",
		//NatsPass:       "T0pS3cr3t",

		Port:           os.Getenv("USER_POST_SERVICE_PORT"),
		UserPostDBHost: os.Getenv("USER_POST_DB_HOST"),
		UserPostDBPort: os.Getenv("MONGO_DB_PORT"),
	}
}
