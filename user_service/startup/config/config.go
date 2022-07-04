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
		Port:       "8003",
		UserDBHost: "localhost",
		UserDBPort: "5432",
		UserDBName: "users",
		UserDBUser: "postgres",
		UserDBPass: "admin",
		NatsHost:   "nats",
		NatsPort:   "4222",
		NatsUser:   "ruser",
		NatsPass:   "T0pS3cr3t",

		//Port:       os.Getenv("USER_SERVICE_PORT"),
		//UserDBHost: os.Getenv("USER_DB_HOST"),
		//UserDBPort: os.Getenv("USER_DB_PORT"),
		//UserDBName: os.Getenv("USER_DB_NAME"),
		//UserDBUser: os.Getenv("USER_DB_USER"),
		//UserDBPass: os.Getenv("USER_DB_PASS"),
	}
}
