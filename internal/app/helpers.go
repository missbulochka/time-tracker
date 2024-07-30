package app

import (
	"fmt"
	"log/slog"
	"time-tracker/internal/adapter/storage/postgres"
	"time-tracker/internal/config"
	usermanager "time-tracker/internal/service/userManager"
)

func getStorageURL(PSQLcfg config.PSQLConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		PSQLcfg.UserName,
		PSQLcfg.Password,
		PSQLcfg.Host,
		PSQLcfg.Port,
		PSQLcfg.DBName,
	)
}

func setupServices(
	log *slog.Logger,
	psqlStorage *postgres.Storage,
) *usermanager.UserManager {
    return usermanager.NewService(
		log,
		psqlStorage,
		psqlStorage,
	)
}
