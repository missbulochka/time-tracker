package http

import (
	"context"
	"log/slog"
	userv1 "time-tracker/internal/handler/http/api/v1/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type UseCase interface {
	DeleteUser(context.Context, uint32) error
}

type Router struct {
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{router: chi.NewRouter()}
}

func (r *Router) WithMiddlewares() *Router {
	r.router.Use(middleware.RequestID)
	r.router.Use(middleware.Logger)
	r.router.Use(middleware.Recoverer)
	// cannot be used by the std
	r.router.Use(middleware.URLFormat)

	return r
}

func (r *Router) AddRoutes(log *slog.Logger, useCase UseCase) *Router {
	r.router.Delete("/users/{user_id}", userv1.NewDeleteHandler(log, useCase))

	return r
}

func (r *Router) GetRouter() *chi.Mux {
	return r.router
}
