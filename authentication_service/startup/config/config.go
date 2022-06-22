package config

import "os"

type Config struct {
	Port       string
	UserDBHost string
	UserDBPort string
	UserDBName string
	UserDBUser string
	UserDBPass string
	NatsHost   string
	NatsPort   string
	NatsUser   string
	NatsPass   string
}

func NewConfig() *Config {
	return &Config{
		Port:       os.Getenv("AUTHENTICATION_SERVICE_PORT"), //"8001",
		UserDBHost: os.Getenv("AUTHENTICATION_DB_HOST"),      //"localhost",
		UserDBPort: os.Getenv("AUTHENTICATION_DB_PORT"),      //"5432",
		UserDBName: os.Getenv("AUTHENTICATION_DB_NAME"),      //"authentication",
		UserDBUser: os.Getenv("AUTHENTICATION_DB_USER"),      //"postgres",
		UserDBPass: os.Getenv("AUTHENTICATION_DB_PASS"),      // "ftn",
		NatsHost:   os.Getenv("NATS_HOST"),                   //" nats",
		NatsPort:   os.Getenv("NATS_PORT"),                   //"4222",
		NatsUser:   os.Getenv("NATS_USER"),                   //"ruser",
		NatsPass:   os.Getenv("NATS_PASS"),                   //"T0pS3cr3t",
	}
}
