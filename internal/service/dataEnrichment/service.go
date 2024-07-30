package dataenrichment

import (
	"context"
	"log/slog"
	"time-tracker/internal/entity"
)

type DataGetter interface {
	GetByID(context.Context, int32, int32) (*entity.User, error)
}

type UserInfo struct {
	log        *slog.Logger
	dataGetter DataGetter
}

func NewService(
	log *slog.Logger,
	dataGetter DataGetter,
) *UserInfo {
	return &UserInfo{
		log:        log,
		dataGetter: dataGetter,
	}
}
