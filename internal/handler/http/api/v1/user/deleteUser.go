package userv1

import (
	"context"
	"log/slog"
	"net/http"
	"time-tracker/internal/entity"
	"time-tracker/internal/entity/validator"
	"time-tracker/internal/handler/http/api/v1/user/models"
	resp "time-tracker/internal/lib/api/response"
	"time-tracker/internal/lib/logger/sl"
	"time-tracker/internal/lib/random"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

const aliasLength = 6

type UserDeleter interface {
	DeleteUser(context.Context, uint32) error
}

func New(log *slog.Logger, userDeleter UserDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "userv1.DeleteUser.New"
		log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req entity.IDRequest
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

		alias := req.Alias
		if len(alias) == 0 {
			alias = random.NewRandomString(aliasLength)
		}

		err := userDeleter.DeleteUser(context.Background(), req.UID)
		if err != nil {
			// TODO: реализация

			return
		}

		render.JSON(w, r, models.Response{
			Response: resp.OK(),
			Alias:    alias,
		})
	}
}
