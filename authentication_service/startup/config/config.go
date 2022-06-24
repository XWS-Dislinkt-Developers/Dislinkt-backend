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
		Port:       "8001",           //os.Getenv("AUTHENTICATION_SERVICE_PORT"),
		UserDBHost: "localhost",      //os.Getenv("AUTHENTICATION_DB_HOST"),
		UserDBPort: "5432",           //os.Getenv("AUTHENTICATION_DB_PORT"),
		UserDBName: "authentication", //os.Getenv("AUTHENTICATION_DB_NAME"),
		UserDBUser: "postgres",       //os.Getenv("AUTHENTICATION_DB_USER"),
		UserDBPass: "ftn",            //os.Getenv("AUTHENTICATION_DB_PASS"),
		NatsHost:   " nats",          //os.Getenv("NATS_HOST"),
		NatsPort:   "4222",           //os.Getenv("NATS_PORT"),
		NatsUser:   "ruser",          //os.Getenv("NATS_USER"),
		NatsPass:   "T0pS3cr3t",      //os.Getenv("NATS_PASS"),
	}
}
