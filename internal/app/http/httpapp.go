package httpapp

import (
	"fmt"
	"log/slog"
	"net/http"
)

type App struct {
	log  *slog.Logger
	addr string
}

func New(
	log *slog.Logger,
	server string,
	port string,
) *App {
	return &App{
		log:  log,
		addr: fmt.Sprintf("%s:%s", server, port),
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "httpapp.Start"

	log := a.log.With(slog.String("op", op))

	r := http.NewServeMux()
	a.RegisterRoutes(r)

	srv := &http.Server{
		Addr:    a.addr,
		Handler: r,
	}

	log.Info("http server is running")

	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
