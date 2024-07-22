package app

import (
	"log/slog"
	httpapp "time-tracker/internal/app/http"
)

type App struct {
	HTTPSrv *httpapp.App
}

func New(
	log *slog.Logger,
	httpServer string,
	httpPort string,
) *App {
	httpapp := httpapp.New(log, httpServer, httpPort)
	return &App{HTTPSrv: httpapp}
}
