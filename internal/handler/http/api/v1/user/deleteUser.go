package userv1

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time-tracker/internal/adapter/storage"
	"time-tracker/internal/entity"
	"time-tracker/internal/entity/validator"
	"time-tracker/internal/handler/http/api/v1/models"
	resp "time-tracker/internal/lib/api/response"
	"time-tracker/internal/lib/logger/sl"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type UserDeleter interface {
	DeleteUser(context.Context, uint32) error
}

func NewDeleteHandler(log *slog.Logger, userDeleter UserDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "userv1.DeleteUser.New"
		log.With(
			"op", op,
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req entity.UIDRequest
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		log.Debug("request body decoded", slog.Any("request", req))

		if err := validator.ValidateStruct(req); err != nil {
			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		err := userDeleter.DeleteUser(context.Background(), req.UID)
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Info("user does not exist", slog.Any("UID", req.UID))

			render.JSON(w, r, resp.Error("user does not exist"))

			return
		}
		if err != nil {
			render.JSON(w, r, resp.Error("failed to delete user"))

			return
		}

		render.JSON(w, r, models.Response{
			Response: resp.OK(),
		})
	}
}
