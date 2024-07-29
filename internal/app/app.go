package app

import (
	"fmt"
	"log/slog"
	"time-tracker/internal/adapter/storage/postgres"
	"time-tracker/internal/config"
	httpapp "time-tracker/internal/handler/http"
)

type App struct {
	log             *slog.Logger
	HTTPSrv         *httpapp.Server
	psqlDatabaseURL string
	psqlDB          *postgres.Storage
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	psqlDatabaseURL := getStorageURL(cfg.PSQLcfg)
	psqlStorage, err := postgres.New(log, psqlDatabaseURL)
	if err != nil {
		panic(err)
	}

	httpapp := httpapp.New(
		log,
		cfg.HTTPcfg.HTTPServer,
		cfg.HTTPcfg.HTTPPort,
	)

	// TODO: remove it
	fmt.Println(psqlStorage)

	return &App{
		log:             log,
		HTTPSrv:         httpapp,
		psqlDatabaseURL: psqlDatabaseURL,
		psqlDB:          psqlStorage,
	}
}
