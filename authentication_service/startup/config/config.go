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
		Port:       os.Getenv("AUTHENTICATION_SERVICE_PORT"), // "8001",           //os.Getenv("AUTHENTICATION_SERVICE_PORT"),
		UserDBHost: os.Getenv("AUTHENTICATION_DB_HOST"),      //"localhost",      //os.Getenv("AUTHENTICATION_DB_HOST"),
		UserDBPort: os.Getenv("AUTHENTICATION_DB_PORT"),      //"5432",           //os.Getenv("AUTHENTICATION_DB_PORT"),
		UserDBName: os.Getenv("AUTHENTICATION_DB_NAME"),      //"authentication", //os.Getenv("AUTHENTICATION_DB_NAME"),
		UserDBUser: os.Getenv("AUTHENTICATION_DB_USER"),      //"postgres",       //os.Getenv("AUTHENTICATION_DB_USER"),
		UserDBPass: os.Getenv("AUTHENTICATION_DB_PASS"),      //"admin",          //,os.Getenv("AUTHENTICATION_DB_PASS"),
		NatsHost:   os.Getenv("NATS_HOST"),                   //"nats",           // os.Getenv("NATS_HOST"),
		NatsPort:   os.Getenv("NATS_PORT"),                   //"4222",           //os.Getenv("NATS_PORT"),
		NatsUser:   os.Getenv("NATS_USER"),                   //"ruser",          //os.Getenv("NATS_USER"),
		NatsPass:   os.Getenv("NATS_PASS"),                   //"T0pS3cr3t",      //os.Getenv("NATS_PASS"),
	}
}
