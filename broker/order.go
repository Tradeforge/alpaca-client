package broker

import (
	"context"
	"net/http"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
)

const (
	EstimateOrderPath = "/v1/trading/accounts/:account_id/orders/estimation"
	CreateOrderPath   = "/v1/trading/accounts/:account_id/orders"
	CancelOrderPath   = "/v1/trading/accounts/:account_id/orders/:order_id"
	GetOrderPath      = "/v1/trading/accounts/:account_id/orders/:order_id"
	ListOrdersPath    = "/v1/trading/accounts/:account_id/orders"
)

type OrderClient struct {
	*client.Client
}

func (oc *OrderClient) EstimateOrder(ctx context.Context, params model.CreateOrderParams, data *model.CreateOrderRequest, opts ...model.RequestOption) (*model.CreateOrderResponse, error) {
	res := &model.CreateOrderResponse{}
	err := oc.Call(ctx, http.MethodPost, EstimateOrderPath, params, res, append(opts, model.Body(data))...)
	return res, err
}

func (oc *OrderClient) CreateOrder(ctx context.Context, params model.CreateOrderParams, data *model.CreateOrderRequest, opts ...model.RequestOption) (*model.CreateOrderResponse, error) {
	res := &model.CreateOrderResponse{}
	err := oc.Call(ctx, http.MethodPost, CreateOrderPath, params, res, append(opts, model.Body(data))...)
	return res, err
}

func (oc *OrderClient) CancelOrder(ctx context.Context, params model.CancelOrderParams, opts ...model.RequestOption) error {
	return oc.Call(ctx, http.MethodDelete, CancelOrderPath, params, nil, opts...)
}

func (oc *OrderClient) ListOrders(ctx context.Context, params model.ListOrdersParams, opts ...model.RequestOption) (model.ListOrdersResponse, error) {
	res := model.ListOrdersResponse{}
	err := oc.Call(ctx, http.MethodGet, ListOrdersPath, params, &res, opts...)
	return res, err
}

func (oc *OrderClient) GetOrder(ctx context.Context, params model.GetOrderParams, opts ...model.RequestOption) (*model.GetOrderResponse, error) {
	res := &model.GetOrderResponse{}
	err := oc.Call(ctx, http.MethodGet, GetOrderPath, params, res, opts...)
	return res, err
}
