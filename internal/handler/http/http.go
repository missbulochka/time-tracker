package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	log  *slog.Logger
	addr string
	srv  *http.Server
}

func New(
	log *slog.Logger,
	server string,
	port string,
) *Server {
	addr := fmt.Sprintf("%s:%s", server, port)
	return &Server{
		log:  log,
		addr: addr,
		srv: &http.Server{
			Addr: addr,
		},
	}
}

func (s *Server) RegisterRouts(mux *chi.Mux) {
	s.srv.Handler = mux
}

func (s *Server) Start() error {
	const op = "http.Start"
	log := s.log.With(slog.String("op", op))

	if s.srv.Handler == nil {
		log.Debug("no routes have registered")
		return fmt.Errorf("no routes have registered")
	}

	log.Info("http server is running")

	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) {
	const op = "http.Stop"
	s.log.With(slog.String("op", op)).Info("stopping http server")

	s.srv.Shutdown(ctx)
}
