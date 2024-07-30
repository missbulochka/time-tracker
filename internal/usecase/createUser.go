package usecase

import (
	"context"
	"time-tracker/internal/entity"
	"time-tracker/internal/lib/logger/sl"
)

func (uc *UseCase) CreateUser(
	ctx context.Context,
	passport *entity.Passport,
	passportSerie, passportNumber int32,
) error {
	const op = "usecase.CreateUser"
	uc.log.With("op", op)

	uc.log.Info("creating user")
	userInfo, err := uc.dataEnrichment.GetUserInfo(ctx, passportSerie, passportNumber)
	if err != nil {
		uc.log.Error("can't get user info", sl.Err(err))
	}

	user := &entity.User{
		Passport: *passport,
		Info:     *userInfo,
	}

	if err := uc.userSaver.AddUser(ctx, user); err != nil {
		uc.log.Error("user addition failed", sl.Err(err))
		return err
	}

	uc.log.Info("user successfully created")
	return nil
}
