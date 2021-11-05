package scaleapi

import (
	"net/http"
)

type Client struct {
	apiKey     string
	httpClient http.Client
}

func Initialize(apiKey string) (Client, error) {
	httpClient := http.Client{}

	client := Client{
		apiKey:     apiKey,
		httpClient: httpClient,
	}

	return client, nil
}
