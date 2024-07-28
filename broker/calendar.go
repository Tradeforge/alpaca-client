package broker

import (
	"context"
	"net/http"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
)

const (
	GetCalendarPath = "/v1/calendar"
)

type MarketClient struct {
	*client.Client
}

func (cc *MarketClient) GetCalendar(ctx context.Context, params *model.GetCalendarParams, opts ...model.RequestOption) (*model.GetCalendarResponse, error) {
	res := &model.GetCalendarResponse{}
	err := cc.Call(ctx, http.MethodGet, GetCalendarPath, params, res, opts...)
	return res, err
}
