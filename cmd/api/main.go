package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/vkzrx/mined/api"
)

func main() {
	ctx := context.Background()

	api := api.New(ctx)
	api.MountHandlers()

	server := &http.Server{Addr: ":" + api.Config.Port, Handler: api.Router}

	go func() {
		log.Printf("Server listening on port %s", api.Config.Port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	nctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stop()
	<-nctx.Done()
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("Shutdown")
}
