package main

import (
    "log"
    "github.com/Qiryl/traffic-control/internal/server";
    "github.com/Qiryl/traffic-control/config";
)

func main() {
	configFile, err := config.LoadConfig("./config/config-local")
	if err != nil {
		log.Fatalf("Load Config: %v", err)
	}

	config, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("Parse Config: %v", err)
	}

	s := server.NewServer(config)
    if err := s.Start(); err != nil {
        log.Fatalf("Server Start: %v", err)
    }
}
