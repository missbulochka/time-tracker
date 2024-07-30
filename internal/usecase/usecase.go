package usecase

import (
	"context"
	"log/slog"
	"time-tracker/internal/entity"
)

type UserSaver interface {
	AddUser(context.Context, *entity.User) error
}

type UserDeleter interface {
	DeleteUser(context.Context, uint32) error
}

type UserGetter interface {
}

type DataEnrichment interface {
	GetUserInfo(context.Context, int32, int32) (*entity.UserInfo, error)
}

type TaskManager interface {
}

type UseCase struct {
	log            *slog.Logger
	userSaver      UserSaver
	userDeleter    UserDeleter
	userGetter     UserGetter
	dataEnrichment DataEnrichment
	taskManager    TaskManager
}

func NewUseCase(
	log *slog.Logger,
	userSaver UserSaver,
	userDeleter UserDeleter,
	userGetter UserGetter,
	dataEnrichment DataEnrichment,
	taskManager TaskManager,
) *UseCase {
	return &UseCase{
		log:            log,
		userSaver:      userSaver,
		userDeleter:    userDeleter,
		userGetter:     userGetter,
		dataEnrichment: dataEnrichment,
		taskManager:    taskManager,
	}
}
