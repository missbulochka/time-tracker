package main

import (
	"log/slog"
	"os"
	"time-tracker/internal/app"
	"time-tracker/internal/config"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	cfg, err := config.LoadCfg()
	if err != nil {
		panic("time-tracker: failed to read config: " + err.Error())
	}

	log := setupLogger(cfg.Env)

	log.Info("starting time-tracker")

	application := app.New(log, cfg.HTTPcfg.HTTPServer, cfg.HTTPcfg.HTTPPort)
	application.HTTPSrv.MustRun()

	// TODO: безопасное окончание программы
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envDev:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
