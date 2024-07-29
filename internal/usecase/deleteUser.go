package usecase

import "context"

func (uc *UseCase) DeleteUser(ctx context.Context, uid uint32) error {
	const op = "usecase.DeleteUser"
	uc.log.With("op", op)

	uc.log.Info("deleting user")
	if err := uc.userDeleter.DeleteUser(ctx, uid); err != nil {
		uc.log.Info("user deletion failed")
		return err
	}

	uc.log.Info("user successfully deleted")
	return nil
}
