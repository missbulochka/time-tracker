package usermanager

import (
	"context"
	"log/slog"
	"time-tracker/internal/entity"
)

type UserSaver interface {
	AddUser(context.Context, *entity.User) error
	// UpdateUser(...) ...
}

type UserDeleter interface {
	DeleteUser(context.Context, uint32) error
}

type UserManager struct {
	log         *slog.Logger
	userSaver   UserSaver
	userDeleter UserDeleter
}

func NewService(
	log *slog.Logger,
	userSaver UserSaver,
	userDeleter UserDeleter,
) *UserManager {
	return &UserManager{
		log:         log,
		userSaver:   userSaver,
		userDeleter: userDeleter,
	}
}
