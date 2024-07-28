package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	Router *chi.Mux
}

func NewRouter() *Router {
	return &Router{Router: chi.NewRouter()}
}

func (r *Router) WithMiddlewares() *Router {
	r.Router.Use(middleware.RequestID)
	r.Router.Use(middleware.Logger)
	r.Router.Use(middleware.Recoverer)
	// cannot be used by the std
	r.Router.Use(middleware.URLFormat)

	return r
}

func (r *Router) WithMethods() *Router {
	// r.Delete("/users/del/", userv1.New(log, useCase))

	return r
}
