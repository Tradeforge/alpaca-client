// Package market defines a REST client for the Alpaca Market Data API.
package market

import (
	"go.tradeforge.dev/alpaca/client"

	"go.uber.org/zap"
)

const (
	apiKeyHeader    = "APCA-API-KEY-ID"
	apiSecretHeader = "APCA-API-SECRET-KEY"
)

// Client defines a client for the Alpaca Broker API.
type Client struct {
	client.Client

	StocksClient
}

// NewClient returns a new client with the specified API key and config.
func NewClient(
	apiURL string,
	apiKey string,
	apiSecret string,
	logger *zap.Logger,
) *Client {
	return newClient(apiURL, apiKey, apiSecret, logger)
}

func newClient(
	apiURL string,
	apiKey string,
	apiSecret string,
	logger *zap.Logger,
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
		Client:       *c,
		StocksClient: StocksClient{Client: *c},
	}
}
