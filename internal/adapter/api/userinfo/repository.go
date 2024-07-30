package userinfo

import "net/http"

type Repository struct {
	client     *http.Client
	clientAddr string
}

func NewRepository(addr string) *Repository {
	return &Repository{
		client:     &http.Client{},
		clientAddr: addr,
	}
}
