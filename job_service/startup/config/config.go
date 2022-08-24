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
		Port:                 "8005",
		UserConnectionDBHost: "localhost",
		UserConnectionDBPort: "27017",
		NatsHost:             "nats",
		NatsPort:             "4222",
		NatsUser:             "ruser",
		NatsPass:             "T0pS3cr3t",

		//Port:                 os.Getenv("JOB_SERVICE_PORT"),
		//UserConnectionDBHost: os.Getenv("JOB_SERVICE_DB_HOST"),
		//UserConnectionDBPort: os.Getenv("MONGO_DB_PORT"),
	}
}
