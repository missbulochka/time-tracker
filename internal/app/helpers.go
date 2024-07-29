package app

import (
	"fmt"
	"time-tracker/internal/config"
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
