package broker

import (
	"context"
	"encoding/json"
	"fmt"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
)

const (
	GetTradeEventsPath    = "/v2beta1/events/trades"
	GetTransferEventsPath = "/v1/events/transfers/status"
)

// EventClient defines a client for the Alpaca Broker Event API.
type EventClient struct {
	*client.Client
}

type TradeEventHandler func(ctx context.Context, event *model.TradeEvent) error

func (c *EventClient) WatchTradeEvents(ctx context.Context, params *model.WatchParams, handler TradeEventHandler, opts ...model.RequestOption) error {
	return c.Listen(
		ctx,
		GetTradeEventsPath,
		params,
		wrapTradeEventHandler(handler),
		opts...,
	)
}

type TransferEventHandler func(ctx context.Context, event *model.TransferEvent) error

func (c *EventClient) SubscribeToTransferEvents(ctx context.Context, params *model.WatchParams, handler TransferEventHandler, opts ...model.RequestOption) error {
	return c.Listen(
		ctx,
		GetTransferEventsPath,
		params,
		wrapTransferEventHandler(handler),
		opts...,
	)
}

func wrapTradeEventHandler(handler TradeEventHandler) client.EventStreamHandler {
	return func(ctx context.Context, event client.Event) error {
		e, err := parseTradeEvent(event)
		if err != nil {
			return fmt.Errorf("parsing trade event: %w", err)
		}
		return handler(ctx, e)
	}
}

func parseTradeEvent(event client.Event) (*model.TradeEvent, error) {
	e := model.TradeEvent{}
	if err := json.Unmarshal(event.GetData(), &e); err != nil {
		return nil, fmt.Errorf("unmarshalling trade event: %w", err)
	}
	return &e, nil
}

func wrapTransferEventHandler(handler TransferEventHandler) client.EventStreamHandler {
	return func(ctx context.Context, event client.Event) error {
		e, err := parseTransferEvent(event)
		if err != nil {
			return fmt.Errorf("parsing transfer event: %w", err)
		}
		return handler(ctx, e)
	}
}

func parseTransferEvent(event client.Event) (*model.TransferEvent, error) {
	e := model.TransferEvent{}
	if err := json.Unmarshal(event.GetData(), &e); err != nil {
		return nil, fmt.Errorf("unmarshalling transfer event: %w", err)
	}
	return &e, nil
}
