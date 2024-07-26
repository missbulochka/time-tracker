package httpapp

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func SetupRoute() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// cannot be used by the std
	router.Use(middleware.URLFormat)

	RegisterRoutes(router)

	return router
}

func RegisterRoutes(r *chi.Mux) {
	// TODO:получение данных пользователей

	// TODO:получение трудозатрат по пользователю за период

	// Начало отчета
}
