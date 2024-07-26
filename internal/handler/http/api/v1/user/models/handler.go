package models

import resp "time-tracker/internal/lib/api/response"

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}
