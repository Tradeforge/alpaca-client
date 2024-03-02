package market

import (
	"context"
	"net/http"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
)

const (
	GetHistoricalBarsPath = "/v2/stocks/bars"
)

// StocksClient is a client for the stocks API.
type StocksClient struct {
	client.Client
}

func (ac *StocksClient) GetHistoricalBars(ctx context.Context, params *model.GetHistoricalBarsParams, opts ...model.RequestOption) (*model.GetHistoricalBarsResponse, error) {
	res := &model.GetHistoricalBarsResponse{}
	err := ac.Call(ctx, http.MethodGet, GetHistoricalBarsPath, params, res)
	return res, err
}
