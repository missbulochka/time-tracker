package httpapp

import (
	"log/slog"
	userv1 "time-tracker/internal/handler/http/api/v1/user"
	"time-tracker/internal/usecase"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func SetupRoute(log *slog.Logger, useCase *usecase.UseCase) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// cannot be used by the std
	router.Use(middleware.URLFormat)

	RegisterRoutes(log, router, useCase)

	return router
}

func RegisterRoutes(log *slog.Logger, r *chi.Mux, useCase *usecase.UseCase) {
	r.Delete("/users/del/", userv1.New(log, useCase))
}
