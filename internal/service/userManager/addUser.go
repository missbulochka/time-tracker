package usermanager

import (
	"context"
	"time-tracker/internal/entity"
	"time-tracker/internal/lib/logger/sl"
)

func (um *UserManager) AddUser(ctx context.Context, user entity.User) error {
	const op = "usermanager.CreateUser"
	um.log.With("op", op)

	um.log.Debug("adding new user to database")
	if err := um.userSaver.AddUser(ctx, user); err != nil {
		um.log.Error("user addition failed", sl.Err(err))

		return err
	}
	um.log.Debug("user successfully added to database")

	return nil
}
