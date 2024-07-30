package userinfo

import (
	"context"
	"time-tracker/internal/entity"
)

const infoEndpoint = "/info"

func (r *Repository) GetByID(context.Context, int32, int32) (*entity.User, error) {
	return nil, nil
}

// func (r *Repository) GetByID(
// 	ctx context.Context,
// 	passportSerie, passportNumber int32,
// ) (*entity.User, error) {
// 	const op = "userinfo.GetByID"

// 	url := r.clientAddr +
// 		infoEndpoint +
// 		fmt.Sprintf("?passportSerie=%s&passportNumber=%s",
// 			string(passportSerie),
// 			string(passportNumber),
// 		)

// 	req, err := http.NewRequestWithContext(
// 		ctx,
// 		http.MethodGet,
// 		url,
// 		nil,
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", op, err)
// 	}

// 	resp, err := r.client.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", op, err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("%s: %d", op, resp.StatusCode)
// 	}

// 	var user entity.User
// 	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
// 		return nil, fmt.Errorf("%s: декодирование ответа: %w", op, err)
// 	}

// 	return &user, nil
// }
