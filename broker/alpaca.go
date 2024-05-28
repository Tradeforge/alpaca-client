// Package broker defines a REST client for the Alpaca Broker API.
package broker

import (
	"log/slog"

	"go.tradeforge.dev/alpaca/client"
)

// Client defines a client for the Alpaca Broker API.
type Client struct {
	*client.Client

	AccountClient
	EventClient
	OrderClient
}

// NewClient returns a new client with the specified API key and config.
func NewClient(
	apiURL string,
	apiKey string,
	apiSecret string,
	reader client.EventReader,
	logger *slog.Logger,
) *Client {
	return newClient(apiURL, apiKey, apiSecret, reader, logger)
}

func newClient(
	apiURL string,
	apiKey string,
	apiSecret string,
	reader client.EventReader,
	logger *slog.Logger,
) *Client {
	c := client.New(
		apiURL,
		reader,
		logger,
	)
	c.SetBasicAuth(apiKey, apiSecret)

	return &Client{
		Client:        c,
		AccountClient: AccountClient{Client: c},
		EventClient:   EventClient{Client: c},
		OrderClient:   OrderClient{Client: c},
	}
}
