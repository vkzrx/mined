package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	Router  *chi.Mux
	Config  *Config
	Context context.Context
}

func NewServer(ctx context.Context) (*Server, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	server := &Server{
		Router:  chi.NewRouter(),
		Config:  config,
		Context: ctx,
	}
	return server, nil
}

func (s *Server) MountHandlers() error {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Heartbeat("/_health"))
	s.Router.Use(cors.Handler(s.Config.Cors))
	s.Router.Use(middleware.Recoverer)

	vm := newVMApi(s.Context)
	mc, err := newMinecraftApi(s.Context)
	if err != nil {
		return err
	}

	s.Router.Mount("/vms", vm.Router)
	s.Router.Mount("/minecraft", mc.Router)

	return nil
}

func (s *Server) Launch() error {
	server := &http.Server{Addr: ":" + s.Config.Port, Handler: s.Router}

	go func() {
		log.Printf("Server listening on port %s", s.Config.Port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	nctx, stop := signal.NotifyContext(s.Context, os.Interrupt, os.Kill)
	defer stop()
	<-nctx.Done()
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(s.Context, 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return err
	}
	log.Println("Server shutdown")

	return nil
}
