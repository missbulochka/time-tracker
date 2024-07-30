package app

import (
	"log/slog"
	"time-tracker/internal/adapter/api/userinfo"
	"time-tracker/internal/adapter/storage/postgres"
	"time-tracker/internal/config"
	httpapp "time-tracker/internal/handler/http"
	"time-tracker/internal/usecase"
)

type App struct {
	log             *slog.Logger
	HTTPSrv         *httpapp.Server
	psqlDatabaseURL string
	psqlDB          *postgres.Storage
	useCase         *usecase.UseCase
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

	userInfoAPI := userinfo.NewRepository(cfg.ExtAPI.UserInfoAPI)

	userManagerServ,
		dataEnrichmentServ := setupServices(log, psqlStorage, userInfoAPI)

	appUseCase := usecase.NewUseCase(
		log,
		userManagerServ,
		userManagerServ,
		// TODO: change structs
		userManagerServ,
		dataEnrichmentServ,
		userManagerServ,
	)

	httpapp := httpapp.New(
		log,
		cfg.HTTPcfg.HTTPServer,
		cfg.HTTPcfg.HTTPPort,
	)

	return &App{
		log:             log,
		HTTPSrv:         httpapp,
		psqlDatabaseURL: psqlDatabaseURL,
		psqlDB:          psqlStorage,
		useCase:         appUseCase,
	}
}
