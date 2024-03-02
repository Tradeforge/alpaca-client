package model

// See https://github.com/alpacahq/alpaca-docs/blob/master/content/api-references/broker-api/events.md.

import (
	"time"
)

type WatchParams struct {
	Since   string `query:"since,omitempty"`
	Until   string `query:"until,omitempty"`
	SinceID string `query:"since_id,omitempty"`
	UntilID string `query:"until_id,omitempty"`
}

// Example of a trade event:
//
//	{
//	    "account_id": "aa4439c3-cf7d-4251-8689-a575a169d6d3",
//	    "at": "2023-10-13T13:28:58.387652Z",
//	    "event_id": "01HCMKKNRK7S5C1JYP50QGDECQ",
//	    "event": "new",
//	    "timestamp": "2023-10-13T13:28:58.37957033Z",
//	    "order": {
//	        "id": "bb2403bc-88ec-430b-b41c-f9ee80c8f0e1",
//	        "client_order_id": "508789e5-cea3-4235-b546-6c62ff92bd79",
//	        "created_at": "2023-10-13T13:28:58.361530031Z",
//	        "updated_at": "2023-10-13T13:28:58.386058029Z",
//	        "submitted_at": "2023-10-13T13:28:58.360070731Z",
//	        "filled_at": null,
//	        "expired_at": null,
//	        "cancel_requested_at": null,
//	        "canceled_at": null,
//	        "failed_at": null,
//	        "replaced_at": null,
//	        "replaced_by": null,
//	        "replaces": null,
//	        "asset_id": "b0b6dd9d-8b9b-48a9-ba46-b9d54906e415",
//	        "symbol": "AAPL",
//	        "asset_class": "us_equity",
//	        "notional": "10",
//	        "qty": null,
//	        "filled_qty": "0",
//	        "filled_avg_price": null,
//	        "order_class": "",
//	        "order_type": "market",
//	        "type": "market",
//	        "side": "buy",
//	        "time_in_force": "day",
//	        "limit_price": null,
//	        "stop_price": null,
//	        "status": "new",
//	        "extended_hours": false,
//	        "legs": null,
//	        "trail_percent": null,
//	        "trail_price": null,
//	        "hwm": null,
//	        "commission": "0"
//	    },
//	    "execution_id": "7922ab44-5b33-4049-ab9a-0cfd805ba989"
//	}
type TradeEvent struct {
	ID          string         `json:"event_id"`
	AccountID   string         `json:"account_id"`
	ExecutionID string         `json:"execution_id"`
	Event       TradeEventType `json:"event"`
	Order       Order          `json:"order"`
	// The average price per share at which the order was filled.
	Price *string `json:"price"`
	// The amount of shares this Trade order was for.
	Quantity *string `json:"qty"`
	// The size of your total position, after this fill event, in shares. Positive for long positions, negative for short positions.
	PositionQuantity *string   `json:"position_qty"`
	Timestamp        time.Time `json:"timestamp"`
}

// See https://github.com/alpacahq/alpaca-docs/blob/master/content/api-references/broker-api/events.md#trade-events.
type TradeEventType string

const (
	/* Common events. */

	// TradeEventNew is sent when an order has been routed to exchanges for execution.
	TradeEventNew TradeEventType = "new"
	// TradeEventFill is sent when your order has been completely filled.
	TradeEventFill TradeEventType = "fill"
	// TradeEventPartialFill is sent when a number of shares less than the total remaining quantity on your order has been filled.
	TradeEventPartialFill TradeEventType = "partial_fill"
	// TradeEventExpired is sent when an order has reached the end of its lifespan, as determined by the order’s time in force value.
	TradeEventExpired TradeEventType = "expired"
	// TradeEventReplaced is sent when your requested replacement of an order is processed.
	TradeEventReplaced TradeEventType = "replaced"
	// TradeEventDoneForDay is sent when the order is done executing for the day, and will not receive further updates until the next trading day.
	TradeEventDoneForDay TradeEventType = "done_for_day"
	// TradeEventCanceled is sent when your requested cancellation of an order is processed.
	TradeEventCanceled TradeEventType = "canceled"

	/* Rarer events */

	// TradeEventRejected is sent when your order has been rejected.
	TradeEventRejected TradeEventType = "rejected"
	// TradeEventPendingNew is sent when the order has been received by Alpaca and routed to the exchanges, but has not yet been accepted for execution.
	TradeEventPendingNew TradeEventType = "pending_new"
	// TradeEventStopped is sent when your order has been stopped, and a trade is guaranteed for the order, usually at a stated price or better, but has not yet occurred.
	TradeEventStopped TradeEventType = "stopped"
	// TradeEventPendingCancel is sent when the order is awaiting cancellation. Most cancellations will occur without the order entering this state.
	TradeEventPendingCancel TradeEventType = "pending_cancel"
	// TradeEventPendingReplace is sent when the order is awaiting replacement.
	TradeEventPendingReplace TradeEventType = "pending_replace"
	// TradeEventCalculated is sent when the order has been completed for the day - it is either filled or done_for_day - but remaining settlement calculations are still pending.
	TradeEventCalculated TradeEventType = "calculated"
	// TradeEventSuspended is sent when the order has been suspended and is not eligible for trading.
	TradeEventSuspended TradeEventType = "suspended"
	// TradeEventOrderReplaceRejected is sent when the order replace has been rejected.
	TradeEventOrderReplaceRejected TradeEventType = "order_replace_rejected"
	// TradeEventOrderCancelRejected is sent when the order cancel has been rejected.
	TradeEventOrderCancelRejected TradeEventType = "order_cancel_rejected"
)

func (e TradeEventType) String() string {
	return string(e)
}

// TradeEvent represents a transfer status update.
type TransferEvent struct {
	ID         string         `json:"event_id"`
	ULID       string         `json:"event_ulid"`
	AccountID  string         `json:"account_id"`
	TransferID string         `json:"transfer_id"`
	StatusFrom TransferStatus `json:"status_from"`
	StatusTo   TransferStatus `json:"status_to"`
	Timestamp  time.Time      `json:"at"`
}
