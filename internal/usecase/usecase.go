package usecase

import (
	"context"
	"log/slog"
)

type UserSaver interface {
}

type UserDeleter interface {
	DeleteUser(context.Context, uint32) error
}

type UserGetter interface {
}

type DataEnrichment interface {
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
    return &UseCase {
        log: log,
        userSaver: userSaver,
        userDeleter: userDeleter,
        userGetter: userGetter,
        dataEnrichment: dataEnrichment,
        taskManager: taskManager,
    }
}
