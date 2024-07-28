package usecase

import (
	"context"
	"log/slog"
)

type UserSaver interface {
}

type UserDeleter interface {
	DeleteUser(ctx context.Context, UID uint32) error
}

type UserGetter interface {
}

type DataEnrichment interface {
}

type TaskManager interface {
}

type UseCase struct {
	log         *slog.Logger
	userDeleter UserDeleter
	// userSaver      UserSaver
	// userGetter     UserGetter
	// dataEnrichment DataEnrichment
	// taskManager    TaskManager
}

func NewUseCase(
	log *slog.Logger,
	userDeleter UserDeleter,
	// userSaver UserSaver,
	// userGetter UserGetter,
	// dataEnrichment DataEnrichment,
	// taskManager TaskManager,
) *UseCase {
	return &UseCase{
		log:         log,
		userDeleter: userDeleter,
		// userSaver:      userSaver,
		// userGetter:     userGetter,
		// dataEnrichment: dataEnrichment,
		// taskManager:    taskManager,
	}
}
