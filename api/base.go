package api

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	Router  *chi.Mux
	Config  *Config
	Context context.Context
}

func New(ctx context.Context) *Server {
	return &Server{
		Router:  chi.NewRouter(),
		Config:  LoadConfig(),
		Context: ctx,
	}
}

func (s *Server) MountHandlers() {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Heartbeat("/_health"))
	s.Router.Use(cors.Handler(s.Config.Cors))
	s.Router.Use(middleware.Recoverer)

	vm := newVMApi(s.Context)

	s.Router.Mount("/", vm.Router)
}
