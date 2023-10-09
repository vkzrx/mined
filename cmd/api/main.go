package main

import (
	"context"
	"log"

	"github.com/vkzrx/mined/api"
)

func main() {
	ctx := context.Background()

	server, err := api.NewServer(ctx)
	if err != nil {
		log.Fatalf("Server failed to initialize: %v", err)
	}

	if err := server.MountHandlers(); err != nil {
		log.Fatalf("Server failed to mount handlers: %v", err)
	}

	if err := server.Launch(); err != nil {
		log.Fatalf("Server failed to launch: %v", err)
	}
}
