package main

import (
	"fmt"
	"log/slog"
	"os"
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

	fmt.Print(cfg)

	log := setupLogger(cfg.Env)

	fmt.Print(log)

	// TODO:запуск сервера

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
