package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	log.Debug("with config",
		slog.String("env", cfg.Env),
		slog.String("addr", cfg.HTTPcfg.HTTPServer+":"+cfg.HTTPcfg.HTTPPort),
	)

	application := app.New(log, cfg)

	application.MustRunPSQLMigration(cfg.MigrationPATH)

	go func() {
		application.MustRunHTTPServer()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sign := <-stop
	log.Info("shutting down time-tracker service", slog.String("signal", sign.String()))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	application.HTTPSrv.Stop(ctx)
	log.Info("time-tracker service stopped")
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
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
