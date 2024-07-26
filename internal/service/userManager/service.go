package usermanager

import (
	"context"
	"log/slog"
)

type UserSaver interface {
	// TODO: add User
	// TODO: Update User
}

type UserDeleter interface {
	DeleteUser(context.Context, string) error
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
