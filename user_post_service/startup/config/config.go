package config

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
		Port:           "8002",
		UserPostDBHost: "localhost",
		UserPostDBPort: "27017",
		NatsHost:       "nats",
		NatsPort:       "4222",
		NatsUser:       "ruser",
		NatsPass:       "T0pS3cr3t",
	}
}
