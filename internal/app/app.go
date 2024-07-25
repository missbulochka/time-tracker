package app

import (
	"fmt"
	"log/slog"
	"time-tracker/internal/adapter/storage/postgres"
	httpapp "time-tracker/internal/app/http"
	"time-tracker/internal/config"
)

type App struct {
	log     *slog.Logger
	HTTPSrv *httpapp.App
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	psqlDatabaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PSQLcfg.UserName,
		cfg.PSQLcfg.Password,
		cfg.PSQLcfg.Host,
		cfg.PSQLcfg.Port,
		cfg.PSQLcfg.DBName,
	)

	if err := postgres.RunMigrate(log, "file://"+cfg.MigrationPATH, psqlDatabaseURL); err != nil {
		panic(err)
	}
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
		log:     log,
		HTTPSrv: httpapp,
	}
}
