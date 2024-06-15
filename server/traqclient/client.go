package traqclient

import (
	"os"

	"github.com/traPtitech/go-traq"
)

var ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")

type Client struct {
	apiClient *traq.APIClient
}

func New(apiClient *traq.APIClient) *Client {
	return &Client{
		apiClient: apiClient,
	}
}
