package broker

import (
	"context"
	"encoding/json"
	"fmt"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
	"go.tradeforge.dev/alpaca/sse"
)

const (
	GetOrderEventsPath         = "/v2beta1/events/trades"
	GetTransferEventPath       = "/v1/events/transfers/status"
	GetAccountStatusEventsPath = "/v1/events/accounts/status"
)

// EventClient defines a client for the Alpaca Broker Event API.
type EventClient struct {
	*client.Client
}

type AccountStatusUpdateHandler func(ctx context.Context, event *model.AccountStatusUpdateEvent) error

// ListenToAccountStatusUpdateEvents listens to account status update SSE events.
// The handler will be called for each event received.
// This is a blocking call.
func (c *EventClient) ListenToAccountStatusUpdateEvents(ctx context.Context, params model.WatchParams, handler AccountStatusUpdateHandler, opts ...model.RequestOption) error {
	return c.Listen(
		ctx,
		GetAccountStatusEventsPath,
		params,
		wrapAccountStatusUpdateHandler(handler),
		opts...,
	)
}

// SubscribeToAccountStatusUpdateEvents subscribes to account status update SSE events.
// The handler will be called for each event received.
// This is a non-blocking call.
func (c *EventClient) SubscribeToAccountStatusUpdateEvents(ctx context.Context, params model.WatchParams, handler AccountStatusUpdateHandler, opts ...model.RequestOption) error {
	return c.Subscribe(
		ctx,
		GetAccountStatusEventsPath,
		params,
		wrapAccountStatusUpdateHandler(handler),
		opts...,
	)
}

func wrapAccountStatusUpdateHandler(handler AccountStatusUpdateHandler) client.EventStreamHandler {
	return func(ctx context.Context, event *sse.Event) error {
		if event.IsComment() {
			return nil
		}
		e, err := parseAccountStatusUpdateEvent(event)
		if err != nil {
			return fmt.Errorf("parsing account status event: %w", err)
		}
		return handler(ctx, e)
	}
}

func parseAccountStatusUpdateEvent(event *sse.Event) (*model.AccountStatusUpdateEvent, error) {
	e := model.AccountStatusUpdateEvent{}
	if err := json.Unmarshal(event.Data, &e); err != nil {
		return nil, fmt.Errorf("unmarshalling account status event: %w", err)
	}
	return &e, nil
}

type TransferStatusUpdateEventHandler func(ctx context.Context, event *model.TransferStatusUpdateEvent) error

// ListenToTransferEvents listens to transfer status update SSE events.
// The handler will be called for each event received.
// This is a blocking call.
func (c *EventClient) ListenToTransferEvents(ctx context.Context, params model.WatchParams, handler TransferStatusUpdateEventHandler, opts ...model.RequestOption) error {
	return c.Listen(
		ctx,
		GetTransferEventPath,
		params,
		wrapTransferEventHandler(handler),
		opts...,
	)
}

// SubscribeToTransferEvents subscribes to transfer status update SSE events.
// The handler will be called for each event received.
// This is a non-blocking call.
func (c *EventClient) SubscribeToTransferEvents(ctx context.Context, params model.WatchParams, handler TransferStatusUpdateEventHandler, opts ...model.RequestOption) error {
	return c.Subscribe(
		ctx,
		GetTransferEventPath,
		params,
		wrapTransferEventHandler(handler),
		opts...,
	)
}

func wrapTransferEventHandler(handler TransferStatusUpdateEventHandler) client.EventStreamHandler {
	return func(ctx context.Context, event *sse.Event) error {
		if event.IsComment() {
			return nil
		}
		e, err := parseTransferEvent(event)
		if err != nil {
			return fmt.Errorf("parsing transfer event: %w", err)
		}
		return handler(ctx, e)
	}
}

func parseTransferEvent(event *sse.Event) (*model.TransferStatusUpdateEvent, error) {
	e := model.TransferStatusUpdateEvent{}
	if err := json.Unmarshal(event.Data, &e); err != nil {
		return nil, fmt.Errorf("unmarshalling transfer event: %w", err)
	}
	return &e, nil
}

type OrderEventHandler func(ctx context.Context, event *model.OrderEvent) error

// ListenToOrderEvents listens to order SSE events.
// The handler will be called for each event received.
// This is a blocking call.
func (c *EventClient) ListenToOrderEvents(ctx context.Context, params model.WatchParams, handler OrderEventHandler, opts ...model.RequestOption) error {
	return c.Listen(
		ctx,
		GetOrderEventsPath,
		params,
		wrapOrderEventHandler(handler),
		opts...,
	)
}

// SubscribeToOrderEvents subscribes to order SSE events.
// The handler will be called for each event received.
// This is a non-blocking call.
func (c *EventClient) SubscribeToOrderEvents(ctx context.Context, params model.WatchParams, handler OrderEventHandler, opts ...model.RequestOption) error {
	return c.Listen(
		ctx,
		GetOrderEventsPath,
		params,
		wrapOrderEventHandler(handler),
		opts...,
	)
}

func wrapOrderEventHandler(handler OrderEventHandler) client.EventStreamHandler {
	return func(ctx context.Context, event *sse.Event) error {
		if event.IsComment() {
			return nil
		}
		e, err := parseOrderEvent(event)
		if err != nil {
			return fmt.Errorf("parsing order event: %w", err)
		}
		return handler(ctx, e)
	}
}

func parseOrderEvent(event *sse.Event) (*model.OrderEvent, error) {
	e := model.OrderEvent{}
	if err := json.Unmarshal(event.Data, &e); err != nil {
		return nil, fmt.Errorf("unmarshalling order event: %w", err)
	}
	return &e, nil
}
