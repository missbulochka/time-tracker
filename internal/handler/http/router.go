package http

import (
	"log/slog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

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

func (r *Router) WithMethods(log *slog.Logger, r *chi.Mux) *Router {
	// r.Delete("/users/del/", userv1.New(log, useCase))

	return r
}
