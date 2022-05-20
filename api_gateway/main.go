package main

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/startup"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/startup/config"
)

func main() {

	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
