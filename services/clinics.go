package services

import (
	"net/http"
)

type HTTPClient interface {
	Get(string) (*http.Response, error)
	Post(string) (*http.Response, error)
}

type Service struct {
	client *http.Client
}

func NewService(client *http.Client) Service {
	return &Service{
		client: client,
	}
}
