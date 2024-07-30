package usecase

import (
	"context"
	"time-tracker/internal/lib/logger/sl"
)

func (uc *UseCase) DeleteUser(ctx context.Context, uid uint32) error {
	const op = "usecase.DeleteUser"
	uc.log.With("op", op)

	uc.log.Info("deleting user")
	if err := uc.userDeleter.DeleteUser(ctx, uid); err != nil {
		uc.log.Error("user deletion failed", sl.Err(err))
		return err
	}

	uc.log.Info("user successfully deleted")
	return nil
}
