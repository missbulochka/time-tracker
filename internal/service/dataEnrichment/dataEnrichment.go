package dataenrichment

import (
	"context"
	"time-tracker/internal/entity"
	"time-tracker/internal/lib/logger/sl"
)

func (ui *UserInfo) GetUserInfo(
	ctx context.Context,
	passportSerie, passportNumber int32,
) (*entity.User, error) {
	const op = "dataenrichment.GetUserInfo"
	ui.log.With("op", op)

	ui.log.Debug("getting userinfo")
	user, err := ui.dataGetter.GetByID(
		ctx,
		passportSerie,
		passportNumber,
	)
	if err != nil {
		ui.log.Error("getting user info failed", sl.Err(err))
		return nil, err
	}
	ui.log.Debug("user info successfully got")

	return user, nil
}
