package main

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/startup"
	cfg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
