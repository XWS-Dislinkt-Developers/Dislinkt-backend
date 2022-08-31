package config

import "os"

type Config struct {
	Port                 string
	UserConnectionDBHost string
	UserConnectionDBPort string
	NatsHost             string
	NatsPort             string
	NatsUser             string
	NatsPass             string

	Neo4jUri      string
	Neo4jHost     string
	Neo4jPort     string
	Neo4jUsername string
	Neo4jPassword string
}

func NewConfig() *Config {
	return &Config{
		//Port:                 "8004",
		//UserConnectionDBHost: "localhost",
		//UserConnectionDBPort: "27017",
		NatsHost: "nats",
		NatsPort: "4222",
		NatsUser: "ruser",
		NatsPass: "T0pS3cr3t",

		Port:                 os.Getenv("USER_CONNECTION_SERVICE_PORT"),
		UserConnectionDBHost: os.Getenv("USER_CONNECTION_DB_HOST"),
		UserConnectionDBPort: os.Getenv("MONGO_DB_PORT"),

		Neo4jUri:      os.Getenv("NEO4J_URI"),
		Neo4jHost:     os.Getenv("NEO4J_HOST"),
		Neo4jPort:     os.Getenv("NEO4J_PORT"),
		Neo4jUsername: os.Getenv("NEO4J_USERNAME"),
		Neo4jPassword: os.Getenv("NEO4J_PASSWORD"),
	}
}
