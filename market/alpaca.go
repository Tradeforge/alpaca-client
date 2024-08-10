// Package market defines a REST client for the Alpaca Market Data API.
package market

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata/stream"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/util"
)

const (
	//nolint:gosec
	apiKeyHeader = "APCA-API-KEY-ID"
	//nolint:gosec
	apiSecretHeader = "APCA-API-SECRET-KEY"

	defaultReconnectInterval = 5 * time.Second
)

type Config struct {
	BaseURL   string `env:"ALPACA_MARKET_API_URL" validate:"required,url"`
	APIKey    string `env:"ALPACA_MARKET_API_KEY" validate:"required"`
	APISecret string `env:"ALPACA_MARKET_API_SECRET" validate:"required"`
	Stream    StreamConfig
}

type StreamConfig struct {
	BaseURL              string         `env:"ALPACA_MARKET_STREAM_API_URL" validate:"required,url"`
	Feed                 string         `env:"ALPACA_MARKET_STREAM_FEED" validate:"required"`
	ReconnectMaxAttempts *int           `env:"ALPACA_MARKET_STREAM_RECONNECT_MAX_ATTEMPTS,default=5"`
	ReconnectInterval    *time.Duration `env:"ALPACA_MARKET_STREAM_RECONNECT_INTERVAL,default=5s"`
}

// Client defines a client for the Alpaca Broker API.
type Client struct {
	*client.Client

	StocksClient
	NewsClient
}

// NewClient returns a new client with the specified API key and config.
func NewClient(
	config Config,
	logger *slog.Logger,
) *Client {
	c := client.New(
		config.BaseURL,
		logger,
	)
	c.SetHeader(apiKeyHeader, config.APIKey)
	c.SetHeader(apiSecretHeader, config.APISecret)

	streamClient := stream.NewStocksClient(
		config.Stream.Feed,
		stream.WithBaseURL(config.Stream.BaseURL),
		stream.WithCredentials(
			config.APIKey,
			config.APISecret,
		),
		stream.WithReconnectSettings(
			*util.Ternary(
				config.Stream.ReconnectMaxAttempts != nil,
				config.Stream.ReconnectMaxAttempts,
				util.AsPtr(0)),
			*util.Ternary(
				config.Stream.ReconnectInterval != nil,
				config.Stream.ReconnectInterval,
				util.AsPtr(defaultReconnectInterval),
			),
		),
		stream.WithConnectCallback(
			func() {
				logger.Debug("connected to stream",
					slog.String("url", config.Stream.BaseURL),
					slog.String("feed", config.Stream.Feed),
				)
			}),
		stream.WithLogger(wrapLogger(logger)),
	)
	return &Client{
		Client: c,
		StocksClient: StocksClient{
			Client: c,
			stream: streamClient,
			logger: logger,
		},
		NewsClient: NewsClient{Client: c},
	}
}

func wrapLogger(logger *slog.Logger) *defaultLogger {
	return &defaultLogger{logger}
}

type defaultLogger struct {
	*slog.Logger
}

func (d *defaultLogger) Infof(format string, v ...interface{}) {
	d.Info(fmt.Sprintf(format, v...))
}

func (d *defaultLogger) Warnf(format string, v ...interface{}) {
	d.Warn(fmt.Sprintf(format, v...))
}

func (d *defaultLogger) Errorf(format string, v ...interface{}) {
	d.Error(fmt.Sprintf(format, v...))
}
