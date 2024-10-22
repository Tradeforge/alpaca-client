package market

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata/stream"
	"github.com/shopspring/decimal"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
)

const (
	GetLatestQuotesPath   = "/v2/stocks/quotes/latest"
	GetSnapshotsPath      = "/v2/stocks/snapshots"
	GetHistoricalBarsPath = "/v2/stocks/bars"
)

// StocksClient is a client for the stocks API.
type StocksClient struct {
	*client.Client
	stream *stream.StocksClient
	logger *slog.Logger
}

func (sc *StocksClient) GetLatestQuotes(ctx context.Context, params model.GetLatestQuotesParams, opts ...model.RequestOption) (*model.GetLatestQuotesResponse, error) {
	res := &model.GetLatestQuotesResponse{}
	err := sc.Call(ctx, http.MethodGet, GetLatestQuotesPath, params, res, opts...)
	return res, err
}

func (sc *StocksClient) GetSnapshots(ctx context.Context, params model.GetSnapshotsParams, opts ...model.RequestOption) (*model.GetSnapshotsResponse, error) {
	res := map[string]model.Snapshot{}
	err := sc.Call(ctx, http.MethodGet, GetSnapshotsPath, params, &res, opts...)
	return &model.GetSnapshotsResponse{Snapshots: res}, err
}

func (sc *StocksClient) GetHistoricalBars(ctx context.Context, params model.GetHistoricalBarsParams, opts ...model.RequestOption) (*model.GetHistoricalBarsResponse, error) {
	res := &model.GetHistoricalBarsResponse{}
	err := sc.Call(ctx, http.MethodGet, GetHistoricalBarsPath, params, res, opts...)
	return res, err
}

type StockBarUpdateHandler func(context.Context, *model.Bar) error

// SubscribeToBarsEvents subscribes to bar updates for the specified symbols.
// The handler is called for each bar update.
// This is a non-blocking call.
func (sc *StocksClient) SubscribeToBarsEvents(
	ctx context.Context,
	params model.StreamStockUpdatesParams,
	handle StockBarUpdateHandler,
) error {
	if err := sc.stream.Connect(ctx); err != nil {
		return err
	}
	return sc.stream.SubscribeToBars(
		func(bar stream.Bar) {
			if err := handle(ctx, &model.Bar{
				Symbol:                     bar.Symbol,
				Open:                       decimal.NewFromFloat(bar.Open),
				High:                       decimal.NewFromFloat(bar.High),
				Low:                        decimal.NewFromFloat(bar.Low),
				Close:                      decimal.NewFromFloat(bar.Close),
				Volume:                     bar.Volume,
				VolumeWeightedAveragePrice: decimal.NewFromFloat(bar.VWAP),
				Timestamp:                  bar.Timestamp,
			}); err != nil {
				// TODO: We might want to optionally unsubscribe from the stream here.
				sc.logger.Error("handling bar", slog.Any("error", err))
			}
		},
		params.Symbols...,
	)
}
