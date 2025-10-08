package main

import (
	"log"
	"meeting_scheduler/config"
	"meeting_scheduler/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	srv, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err := srv.Run(cfg.ServerPort); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
