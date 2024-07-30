package userinfo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time-tracker/internal/entity"
)

const infoEndpoint = "/info"

func (r *Repository) GetByPassport(
	ctx context.Context,
	passportSerie, passportNumber int32,
) (*entity.UserInfo, error) {
	const op = "userinfo.GetByID"

	url := r.clientAddr +
		infoEndpoint +
		fmt.Sprintf("?passportSerie=%s&passportNumber=%s",
			string(passportSerie),
			string(passportNumber),
		)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %d", op, resp.StatusCode)
	}

	var user entity.UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}
