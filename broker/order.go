package broker

import (
	"context"
	"net/http"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
)

const (
	CreateOrderPath = "/v1/trading/accounts/:id/orders"
)

type OrderClient struct {
	client.Client
}

func (oc *OrderClient) CreateOrder(ctx context.Context, params *model.CreateOrderParams, data *model.CreateOrderRequest, opts ...model.RequestOption) (*model.CreateOrderResponse, error) {
	res := &model.CreateOrderResponse{}
	err := oc.Call(ctx, http.MethodPost, CreateOrderPath, params, res, model.Body(data))
	return res, err
}
