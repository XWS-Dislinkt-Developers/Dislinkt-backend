package config

type Config struct {
	Port                 string
	UserConnectionDBHost string
	UserConnectionDBPort string
	NatsHost             string
	NatsPort             string
	NatsUser             string
	NatsPass             string
}

func NewConfig() *Config {
	return &Config{
		Port:                 "8004",      // os.Getenv("ORDERING_SERVICE_PORT"),
		UserConnectionDBHost: "localhost", // os.Getenv("ORDERING_DB_HOST"),
		UserConnectionDBPort: "27017",     // os.Getenv("ORDERING_DB_PORT"),
		NatsHost:             "nats",      // os.Getenv("NATS_HOST"),
		NatsPort:             "4222",      // os.Getenv("NATS_PORT"),
		NatsUser:             "ruser",     // os.Getenv("NATS_USER"),
		NatsPass:             "T0pS3cr3t", // os.Getenv("NATS_PASS"),
	}
}
