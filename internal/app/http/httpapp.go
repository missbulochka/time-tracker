package httpapp

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time-tracker/internal/usecase"
)

type App struct {
	log  *slog.Logger
	addr string
	srv  *http.Server
}

func New(
	log *slog.Logger,
	server string,
	port string,
	useCase *usecase.UseCase,
) *App {
	router := SetupRoute(log, useCase)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", server, port),
		Handler: router,
	}

	return &App{
		log:  log,
		addr: srv.Addr,
		srv:  srv,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "httpapp.Start"
	log := a.log.With(slog.String("op", op))

	log.Info("http server is running")

	return a.srv.ListenAndServe()

}

func (a *App) Stop(ctx context.Context) {
	const op = "httpapp.Stop"
	a.log.With(slog.String("op", op)).Info("stopping http server")

	a.srv.Shutdown(ctx)
}
