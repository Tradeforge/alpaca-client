package market

import (
	"context"
	"net/http"

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
}

func (ac *StocksClient) GetLatestQuotes(ctx context.Context, params *model.GetLatestQuotesParams, opts ...model.RequestOption) (*model.GetLatestQuotesResponse, error) {
	res := &model.GetLatestQuotesResponse{}
	err := ac.Call(ctx, http.MethodGet, GetLatestQuotesPath, params, res, opts...)
	return res, err
}

func (ac *StocksClient) GetSnapshots(ctx context.Context, params *model.GetSnapshotsParams, opts ...model.RequestOption) (*model.GetSnapshotsResponse, error) {
	res := map[string]model.Snapshot{}
	err := ac.Call(ctx, http.MethodGet, GetSnapshotsPath, params, &res, opts...)
	return &model.GetSnapshotsResponse{Snapshots: res}, err
}

func (ac *StocksClient) GetHistoricalBars(ctx context.Context, params *model.GetHistoricalBarsParams, opts ...model.RequestOption) (*model.GetHistoricalBarsResponse, error) {
	res := &model.GetHistoricalBarsResponse{}
	err := ac.Call(ctx, http.MethodGet, GetHistoricalBarsPath, params, res, opts...)
	return res, err
}
