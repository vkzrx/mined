package main

import (
	"log"

	"github.com/vkzrx/mined/vm"
)

func main() {
	server, err := vm.NewMinecraftServiceServer("5001")
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	if err := server.Launch(); err != nil {
		log.Fatalf("Server failed to launch: %v", err)
	}
}
