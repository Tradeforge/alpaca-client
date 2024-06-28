package broker

import (
	"context"
	"encoding/json"
	"fmt"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
)

const (
	GetTradeEventsPath          = "/v2beta1/events/trades"
	GetTransferStatusEventsPath = "/v1/events/transfers/status"
	GetAccountStatusEventsPath  = "v1/events/accounts/status"
)

// EventClient defines a client for the Alpaca Broker Event API.
type EventClient struct {
	*client.Client
}

type AccountStatusUpdateHandler func(ctx context.Context, event *model.AccountStatusUpdateEvent) error

func (c *EventClient) SubscribeToAccountStatusUpdateEvents(ctx context.Context, params *model.WatchParams, handler AccountStatusUpdateHandler, opts ...model.RequestOption) error {
	return c.Listen(
		ctx,
		GetAccountStatusEventsPath,
		params,
		wrapAccountStatusUpdateHandler(handler),
		opts...,
	)
}

func wrapAccountStatusUpdateHandler(handler AccountStatusUpdateHandler) client.EventStreamHandler {
	return func(ctx context.Context, event client.Event) error {
		e, err := parseAccountStatusUpdateEvent(event)
		if err != nil {
			return fmt.Errorf("parsing account status event: %w", err)
		}
		return handler(ctx, e)
	}
}

func parseAccountStatusUpdateEvent(event client.Event) (*model.AccountStatusUpdateEvent, error) {
	e := model.AccountStatusUpdateEvent{}
	if err := json.Unmarshal(event.GetData(), &e); err != nil {
		return nil, fmt.Errorf("unmarshalling account status event: %w", err)
	}
	return &e, nil
}

type TransferStatusUpdateEventHandler func(ctx context.Context, event *model.TransferStatusUpdateEvent) error

func (c *EventClient) SubscribeToTransferStatusUpdateEvents(ctx context.Context, params *model.WatchParams, handler TransferStatusUpdateEventHandler, opts ...model.RequestOption) error {
	return c.Listen(
		ctx,
		GetTransferStatusEventsPath,
		params,
		wrapTransferEventHandler(handler),
		opts...,
	)
}

func wrapTransferEventHandler(handler TransferStatusUpdateEventHandler) client.EventStreamHandler {
	return func(ctx context.Context, event client.Event) error {
		e, err := parseTransferEvent(event)
		if err != nil {
			return fmt.Errorf("parsing transfer event: %w", err)
		}
		return handler(ctx, e)
	}
}

func parseTransferEvent(event client.Event) (*model.TransferStatusUpdateEvent, error) {
	e := model.TransferStatusUpdateEvent{}
	if err := json.Unmarshal(event.GetData(), &e); err != nil {
		return nil, fmt.Errorf("unmarshalling transfer event: %w", err)
	}
	return &e, nil
}

type TradeEventHandler func(ctx context.Context, event *model.TradeEvent) error

func (c *EventClient) SubscribeToTradeEvents(ctx context.Context, params *model.WatchParams, handler TradeEventHandler, opts ...model.RequestOption) error {
	return c.Listen(
		ctx,
		GetTradeEventsPath,
		params,
		wrapTradeEventHandler(handler),
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
