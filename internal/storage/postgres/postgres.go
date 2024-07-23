package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time-tracker/internal/lib/logger/sl"

	_ "github.com/lib/pq" // init postgresql driver
)

type Storage struct {
	log *slog.Logger
	db  *sql.DB
}

func New(log *slog.Logger, storageSetup string) (*Storage, error) {
	const op = "storage.postgres.New"
	log.With(slog.String("op", op))

	db, err := sql.Open("postgres", storageSetup)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Debug("successfully connected to psql")

	return &Storage{
		log: log,
		db:  db,
	}, nil
}
