package broker

import (
    "context"
    "net/http"

    "go.tradeforge.dev/alpaca/client"
    "go.tradeforge.dev/alpaca/model"
)

const (
    GetOpenPositionBySymbolPath = "/v1/trading/accounts/:account_id/positions/:symbol"
    ListOpenPositionsPath       = "/v1/trading/accounts/:account_id/positions"
)

type TradingClient struct {
    *client.Client
}

func (tc *TradingClient) GetOpenPositionBySymbol(ctx context.Context, params model.GetOpenPositionBySymbolParams, opts ...model.RequestOption) (*model.GetOpenPositionResponse, error) {
    res := &model.GetOpenPositionResponse{}
    err := tc.Call(ctx, http.MethodGet, GetOpenPositionBySymbolPath, params, res, opts...)
    return res, err
}

func (tc *TradingClient) ListOpenPositions(ctx context.Context, params model.ListOpenPositionsParams, opts ...model.RequestOption) ([]model.GetOpenPositionResponse, error) {
    var res []model.GetOpenPositionResponse
    err := tc.Call(ctx, http.MethodGet, ListOpenPositionsPath, params, &res, opts...)
    return res, err
}
