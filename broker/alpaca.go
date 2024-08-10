// Package broker defines a REST client for the Alpaca Broker API.
package broker

import (
	"log/slog"

	"go.tradeforge.dev/alpaca/client"
)

type Config struct {
	BaseURL   string `env:"ALPACA_BROKER_API_URL" validate:"required,url"`
	APIKey    string `env:"ALPACA_BROKER_API_KEY" validate:"required"`
	APISecret string `env:"ALPACA_BROKER_API_SECRET" validate:"required"`
}

// Client defines a client for the Alpaca Broker API.
type Client struct {
	*client.Client

	AccountClient
	EventClient
	OrderClient
	MarketClient
	TradingClient
}

// NewClient returns a new client with the specified API key and config.
func NewClient(
	config Config,
	reader client.EventReader,
	logger *slog.Logger,
) *Client {
	c := client.New(
		config.BaseURL,
		logger,
	)
	c.SetBasicAuth(config.APIKey, config.APISecret)
	return &Client{
		Client:        c,
		AccountClient: AccountClient{Client: c},
		EventClient:   EventClient{Client: c},
		OrderClient:   OrderClient{Client: c},
		MarketClient:  MarketClient{Client: c},
		TradingClient: TradingClient{Client: c},
	}
}
