package market

import (
	"context"
	"net/http"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
)

const (
	GetNewsPath = "/v1beta1/news"
)

type NewsClient struct {
	*client.Client
}

func (nc *NewsClient) GetLatestNews(ctx context.Context, params model.GetNewsParams, opts ...model.RequestOption) (*model.GetNewsResponse, error) {
	res := &model.GetNewsResponse{}
	err := nc.Call(ctx, http.MethodGet, GetNewsPath, params, res, opts...)
	return res, err
}
