package app

import (
	"fmt"
	"log/slog"
	"time-tracker/internal/adapter/api/userinfo"
	"time-tracker/internal/adapter/storage/postgres"
	"time-tracker/internal/config"
	dataenrichment "time-tracker/internal/service/dataEnrichment"
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
	userInfoAPI *userinfo.Repository,
) (
	*usermanager.UserManager,
	*dataenrichment.UserInfo,
) {
	return usermanager.NewService(
			log,
			psqlStorage,
			psqlStorage,
		),
		dataenrichment.NewService(
			log,
			userInfoAPI,
		)
}
