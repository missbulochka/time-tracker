package usermanager

import (
	"context"
	"time-tracker/internal/lib/logger/sl"
)

func (um *UserManager) DeleteUser(ctx context.Context, uid uint32) error {
	const op = "usermanager.DeleteUser"
	um.log.With("op", op)

	um.log.Debug("deleting user from database")
	if err := um.userDeleter.DeleteUser(ctx, uid); err != nil {
		um.log.Error("user deletion failed", sl.Err(err))
		return err
	}
	um.log.Debug("user successfully deleted from database")

	return nil
}
