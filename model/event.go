package model

// See https://github.com/alpacahq/alpaca-docs/blob/master/content/api-references/broker-api/events.md.

import (
	"time"

	"github.com/google/uuid"
)

type WatchParams struct {
	// Since is a string field that represents a date in the format "YYYY-MM-DD". It is used to specify the start date from which to watch events.
	Since string `query:"since,omitempty"`
	// Until s a string field that represents a date in the format "YYYY-MM-DD". It is used to specify the end date until which to watch events.
	Until string `query:"until,omitempty"`
	// SinceID is a string field that is used to specify the ID from which to start watching events.
	SinceID string `query:"since_id,omitempty"`
	// UntilID is a string field that is used to specify the ID until which to watch events.
	UntilID string `query:"until_id,omitempty"`
}

// AccountStatusUpdateEvent represents an account status event.
type AccountStatusUpdateEvent struct {
	EventID             int                        `json:"event_id"`
	EventUlid           string                     `json:"event_ulid"`
	AccountID           uuid.UUID                  `json:"account_id"`
	AccountNumber       string                     `json:"account_number"`
	StatusFrom          AccountStatus              `json:"status_from"`
	StatusTo            AccountStatus              `json:"status_to"`
	Reason              string                     `json:"reason"`
	At                  string                     `json:"at"`
	KYCResults          KYCResults                 `json:"kyc_results"`
	CryptoStatusFrom    string                     `json:"crypto_status_from"`
	CryptoStatusTo      string                     `json:"crypto_status_to"`
	AdminConfigurations AccountAdminConfigurations `json:"admin_configurations"`
	PatternDayTrader    bool                       `json:"pattern_day_trader"`
	AccountBlocked      bool                       `json:"account_blocked"`
	TradingBlocked      bool                       `json:"trading_blocked"`
}

// OrderEvent represents an order event.
type OrderEvent struct {
	ID          string         `json:"event_id"`
	AccountID   uuid.UUID      `json:"account_id"`
	ExecutionID uuid.UUID      `json:"execution_id"`
	Event       OrderEventType `json:"event"`
	Order       Order          `json:"order"`
	// The average price per share at which the order was filled.
	Price *string `json:"price"`
	// The amount of shares this Trade order was for.
	Quantity *string `json:"qty"`
	// The size of your total position, after this fill event, in shares. Positive for long positions, negative for short positions.
	PositionQuantity *string   `json:"position_qty"`
	Timestamp        time.Time `json:"timestamp"`
}

// OrderEventType represents an order event type.
//
// See https://github.com/alpacahq/alpaca-docs/blob/master/content/api-references/broker-api/events.md#trade-events.
type OrderEventType string

const (
	/* Common events. */

	// OrderEventNew is sent when an order has been routed to exchanges for execution.
	OrderEventNew OrderEventType = "new"
	// OrderEventFill is sent when your order has been completely filled.
	OrderEventFill OrderEventType = "fill"
	// OrderEventPartialFill is sent when a number of shares less than the total remaining quantity on your order has been filled.
	OrderEventPartialFill OrderEventType = "partial_fill"
	// OrderEventExpired is sent when an order has reached the end of its lifespan, as determined by the order’s time in force value.
	OrderEventExpired OrderEventType = "expired"
	// OrderEventReplaced is sent when your requested replacement of an order is processed.
	OrderEventReplaced OrderEventType = "replaced"
	// OrderEventDoneForDay is sent when the order is done executing for the day, and will not receive further updates until the next trading day.
	OrderEventDoneForDay OrderEventType = "done_for_day"
	// OrderEventCanceled is sent when your requested cancellation of an order is processed.
	OrderEventCanceled OrderEventType = "canceled"

	/* Rarer events */

	// OrderEventRejected is sent when your order has been rejected.
	OrderEventRejected OrderEventType = "rejected"
	// OrderEventPendingNew is sent when the order has been received by Alpaca and routed to the exchanges, but has not yet been accepted for execution.
	OrderEventPendingNew OrderEventType = "pending_new"
	// OrderEventStopped is sent when your order has been stopped, and a trade is guaranteed for the order, usually at a stated price or better, but has not yet occurred.
	OrderEventStopped OrderEventType = "stopped"
	// OrderEventPendingCancel is sent when the order is awaiting cancellation. Most cancellations will occur without the order entering this state.
	OrderEventPendingCancel OrderEventType = "pending_cancel"
	// OrderEventPendingReplace is sent when the order is awaiting replacement.
	OrderEventPendingReplace OrderEventType = "pending_replace"
	// OrderEventCalculated is sent when the order has been completed for the day - it is either filled or done_for_day - but remaining settlement calculations are still pending.
	OrderEventCalculated OrderEventType = "calculated"
	// OrderEventSuspended is sent when the order has been suspended and is not eligible for trading.
	OrderEventSuspended OrderEventType = "suspended"
	// OrderEventOrderReplaceRejected is sent when the order replace has been rejected.
	OrderEventOrderReplaceRejected OrderEventType = "order_replace_rejected"
	// OrderEventOrderCancelRejected is sent when the order cancel has been rejected.
	OrderEventOrderCancelRejected OrderEventType = "order_cancel_rejected"
)

func (e OrderEventType) String() string {
	return string(e)
}

// TransferStatusUpdateEvent represents a transfer status update.
type TransferStatusUpdateEvent struct {
	ID         string         `json:"event_id"`
	ULID       string         `json:"event_ulid"`
	AccountID  uuid.UUID      `json:"account_id"`
	TransferID uuid.UUID      `json:"transfer_id"`
	StatusFrom TransferStatus `json:"status_from"`
	StatusTo   TransferStatus `json:"status_to"`
	Timestamp  time.Time      `json:"at"`
}

type TransferStatus string

const (
	// TransferStatusQueued represents a transfer that is in queue to be processed.
	TransferStatusQueued TransferStatus = "QUEUED"
	// TransferStatusApprovalPending represents a transfer that is pending approval.
	TransferStatusApprovalPending TransferStatus = "APPROVAL_PENDING"
	// TransferStatusPending represents a transfer that is pending processing.
	TransferStatusPending TransferStatus = "PENDING"
	// TransferStatusSentToClearing represents a transfer that is being processed by the clearing firm.
	TransferStatusSentToClearing TransferStatus = "SENT_TO_CLEARING"
	// TransferStatusRejected represents a transfer that is rejected.
	TransferStatusRejected TransferStatus = "REJECTED"
	// TransferStatusCanceled represents a client initiated transfer cancellation.
	TransferStatusCanceled TransferStatus = "CANCELED"
	// TransferStatusApproved represents a transfer that is approved.
	TransferStatusApproved TransferStatus = "APPROVED"
	// TransferStatusComplete represents a transfer that is completed.
	TransferStatusComplete TransferStatus = "COMPLETE"
	// TransferStatusReturned represents a bank issued ACH return for the transfer.
	TransferStatusReturned TransferStatus = "RETURNED"
)
