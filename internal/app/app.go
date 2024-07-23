package app

import (
	"fmt"
	"log/slog"
	httpapp "time-tracker/internal/app/http"
	"time-tracker/internal/config"
	"time-tracker/internal/storage/postgres"
)

type App struct {
	log     *slog.Logger
	HTTPSrv *httpapp.App
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	httpapp := httpapp.New(
		log,
		cfg.HTTPcfg.HTTPServer,
		cfg.HTTPcfg.HTTPPort,
	)

	psqlStorage, err := postgres.New(
		log,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.PSQLcfg.Host,
			cfg.PSQLcfg.Port,
			cfg.PSQLcfg.UserName,
			cfg.PSQLcfg.Password,
			cfg.PSQLcfg.DBName,
		),
	)
	if err != nil {
		panic(err)
	}

	// TODO: remove it
	fmt.Println(psqlStorage)

	return &App{
		log:     log,
		HTTPSrv: httpapp,
	}
}
