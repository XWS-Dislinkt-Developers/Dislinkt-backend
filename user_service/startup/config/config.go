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
		Port:       "8005",      //os.Getenv("USER_SERVICE_PORT"),
		UserDBHost: "localhost", //os.Getenv("USER_DB_HOST"),
		UserDBPort: "5432",      //os.Getenv("USER_DB_PORT"),
		UserDBName: "users",     //os.Getenv("USER_DB_NAME"),
		UserDBUser: "postgres",  //os.Getenv("USER_DB_USER"),
		UserDBPass: "postgres",  // os.Getenv("USER_DB_PASS"),
		NatsHost:   " nats",     //os.Getenv("NATS_HOST"),
		NatsPort:   "4222",      //os.Getenv("NATS_PORT"),
		NatsUser:   "ruser",     //os.Getenv("NATS_USER"),
		NatsPass:   "T0pS3cr3t", //os.Getenv("NATS_PASS"),
	}
}
