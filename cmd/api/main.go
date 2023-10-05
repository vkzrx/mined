package main

import (
	"context"
	"log"

	"github.com/vkzrx/mined/api"
)

func main() {
	ctx := context.Background()

	server := api.NewServer(ctx)

	server.MountHandlers()

	if err := server.Launch(); err != nil {
		log.Fatalf("Server failed to launch: %v", err)
	}
}
