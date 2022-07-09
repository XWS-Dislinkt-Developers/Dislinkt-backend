package config

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

		Port:       "8001",
		UserDBHost: "localhost",
		UserDBPort: "5432",
		UserDBName: "authentication",
		UserDBUser: "postgres",
		UserDBPass: "postgres",
		NatsHost:   "nats",
		NatsPort:   "4222",
		NatsUser:   "ruser",
		NatsPass:   "T0pS3cr3t",


		/*
			Port:       os.Getenv("AUTHENTICATION_SERVICE_PORT"),
			UserDBHost: os.Getenv("AUTHENTICATION_DB_HOST"),
			UserDBPort: os.Getenv("AUTHENTICATION_DB_PORT"),
			UserDBName: os.Getenv("AUTHENTICATION_DB_NAME"),
			UserDBUser: os.Getenv("AUTHENTICATION_DB_USER"),
			UserDBPass: os.Getenv("AUTHENTICATION_DB_PASS"),
			NatsHost:   os.Getenv("NATS_HOST"),
			NatsPort:   os.Getenv("NATS_PORT"),
			NatsUser:   os.Getenv("NATS_USER"),
			NatsPass:   os.Getenv("NATS_PASS"),


		*/
	}
}
