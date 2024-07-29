package postgres

import (
	"errors"
	"log/slog"
	"time-tracker/internal/lib/logger/sl"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (s *Storage) RunMigrate(sourceURL, databaseURL string) error {
	const op = "postgres.Migrate"
	s.log.With(slog.String("op", op))

	migration, err := migrate.New(sourceURL, databaseURL)

	if err != nil {
		s.log.Error("migration failed", sl.Err(err))
		return err
	}

	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		s.log.Error("migration failed", sl.Err(err))
		return err
	}

	s.log.Debug("migrated successfully")

	return nil
}
