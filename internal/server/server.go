package server

import (
	"auth/internal/config"
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg *config.HTTPServer, router http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         cfg.Addr,
			Handler:      router,
			ReadTimeout:  cfg.Timeout,
			WriteTimeout: cfg.Timeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
