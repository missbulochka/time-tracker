package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time-tracker/internal/adapter/storage"
	"time-tracker/internal/lib/logger/sl"

	_ "github.com/lib/pq" // init postgresql driver
)

type Storage struct {
	log *slog.Logger
	db  *sql.DB
}

func New(log *slog.Logger, databaseURL string) (*Storage, error) {
	const op = "postgres.New"
	log.With(slog.String("op", op))

	db, err := sql.Open("postgres", databaseURL)
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

func (s *Storage) DeleteUser(ctx context.Context, uid uint32) error {
	var exists bool
	err := s.db.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM users WHERE user_id=$1)",
		uid,
	).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return storage.ErrUserNotFound
	}

	_, err = s.db.ExecContext(
		ctx,
		"DELETE FROM users WHERE user_id=$1",
		uid,
	)
	if err != nil {
		return err
	}

	return nil
}
