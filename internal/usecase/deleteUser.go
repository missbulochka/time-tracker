package usecase

import "context"

func (uc *UseCase) DeleteUser(
	ctx context.Context,
	UID string,
) error {
	if err := uc.userDeleter.DeleteUser(ctx, UID); err != nil {
		return err
	}

	return nil
}
