package usermanager

import (
	"context"
	"log/slog"
	"time-tracker/internal/lib/logger/sl"
)

func (um *UserManager) DeleteUser(ctx context.Context, UID string) error {
	const op = "usermanager.DeleteUser"
	um.log.With(slog.String("op", op))

	if err := um.userDeleter.DeleteUser(ctx, UID); err != nil {
		um.log.Error("user deletion failed", sl.Err(err))
		return err
	}

	um.log.Debug("user seccessfully deleted")

	return nil
}
