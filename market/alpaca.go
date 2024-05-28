// Package market defines a REST client for the Alpaca Market Data API.
package market

import (
	"log/slog"

	"go.tradeforge.dev/alpaca/client"
)

const (
	//nolint:gosec
	apiKeyHeader = "APCA-API-KEY-ID"
	//nolint:gosec
	apiSecretHeader = "APCA-API-SECRET-KEY"
)

// Client defines a client for the Alpaca Broker API.
type Client struct {
	*client.Client

	StocksClient
	NewsClient
}

// NewClient returns a new client with the specified API key and config.
func NewClient(
	apiURL string,
	apiKey string,
	apiSecret string,
	logger *slog.Logger,
) *Client {
	return newClient(apiURL, apiKey, apiSecret, logger)
}

func newClient(
	apiURL string,
	apiKey string,
	apiSecret string,
	logger *slog.Logger,
) *Client {
	c := client.New(
		apiURL,
		nil,
		logger,
	)
	c.SetHeader(
		apiKeyHeader,
		apiKey,
	)
	c.SetHeader(
		apiSecretHeader,
		apiSecret,
	)
	return &Client{
		Client:       c,
		StocksClient: StocksClient{Client: c},
		NewsClient:   NewsClient{Client: c},
	}
}
