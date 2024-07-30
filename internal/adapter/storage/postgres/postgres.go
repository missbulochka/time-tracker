package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time-tracker/internal/adapter/storage"
	"time-tracker/internal/entity"
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
	const op = "postgres.DeleteUser"

	if _, err := s.GetUser(ctx, uid); err != nil {
		if err == storage.ErrUserNotFound {
			return err
		}
	}

	_, err := s.db.ExecContext(
		ctx,
		"DELETE FROM users WHERE user_id=$1",
		uid,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) AddUser(
	ctx context.Context,
	user *entity.User,
) error {
	const op = "postgres.AddUser"

	stmt, err := s.db.Prepare(`
	INSERT INTO users(
	    passport_number,
		surname,
		name,
		patronymic,
		adress)
	VALUES($1, $2, $3, $4, $5)
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		user.Passport.PasspotNumber,
		user.Info.Surname,
		user.Info.Name,
		user.Info.Patronymic,
		user.Info.Adress,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetUser(ctx context.Context, uid uint32) (*entity.User, error) {
	const op = "postgres.GetUser"

	var user entity.User
	err := s.db.QueryRowContext(
		ctx,
		`SELECT (
		    passport_number,
			surname,
			name,
			patronymic,
			adress,
		) FROM users WHERE user_id=$1`,
		uid,
	).Scan(&user.Passport.PasspotNumber,
		&user.Info.Surname,
		&user.Info.Name,
		&user.Info.Patronymic,
		&user.Info.Adress,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(user.Passport.PasspotNumber) == 0 {
		return nil, storage.ErrUserNotFound
	}

	return &user, nil
}
