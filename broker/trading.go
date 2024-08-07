package broker

import (
	"context"
	"net/http"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/errors"
	"go.tradeforge.dev/alpaca/model"
)

const (
	GetOpenPositionBySymbolPath = "/v1/trading/accounts/:account_id/positions/:symbol"
)

type TradingClient struct {
	*client.Client
}

func (tc *TradingClient) GetOpenPositionBySymbol(ctx context.Context, params *model.GetOpenPositionBySymbolParams, opts ...model.RequestOption) (*model.GetOpenPositionResponse, error) {
	res := &model.GetOpenPositionResponse{}
	err := tc.Call(ctx, http.MethodGet, GetOpenPositionBySymbolPath, params, res, opts...)
	if responseErr, ok := errors.AsResponseError(err); ok {
		switch responseErr.StatusCode {
		case http.StatusNotFound:
			return nil, errors.NewPositionFoundError()
		}
	}
	return res, err
}
