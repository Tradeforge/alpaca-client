package model

import (
	"time"
)

type GetLatestQuotesParams struct {
	Symbols  string  `query:"symbols,required"`
	Feed     *string `query:"feed,omitempty"`
	Currency *string `query:"currency,omitempty"`
}

type GetLatestQuotesResponse struct {
	Quotes   map[string]Quote `json:"quotes"`
	Currency string           `json:"currency"`
}

type Quote struct {
	AskPrice  float64   `json:"ap"`
	AskSize   uint64    `json:"as"`
	BidPrice  float64   `json:"bp"`
	BidSize   uint64    `json:"bs"`
	Timestamp time.Time `json:"t"`
}

type GetLatestTradesParams struct {
	Symbols  string  `query:"symbols,required"`
	Feed     *string `query:"feed,omitempty"`
	Currency *string `query:"currency,omitempty"`
}

type GetLatestTradesResponse struct {
	Trades   map[string]LatestTrade `json:"trades"`
	Currency string                 `json:"currency"`
}

type LatestTrade struct {
	Price     float64   `json:"p"`
	Size      uint64    `json:"s"`
	Exchange  string    `json:"x"`
	Timestamp time.Time `json:"t"`
}

type GetSnapshotsParams struct {
	Symbols  string  `query:"symbols,required"`
	Feed     *string `query:"feed,omitempty"`
	Currency *string `query:"currency,omitempty"`
}

type GetSnapshotsResponse struct {
	Snapshots map[string]Snapshot
}

type Snapshot struct {
	LatestTrade    LatestTrade `json:"latestTrade"`
	LatestQuote    Quote       `json:"latestQuote"`
	MinBar         Bar         `json:"minuteBar"`
	DayBar         Bar         `json:"dailyBar"`
	PreviousDayBar Bar         `json:"prevDailyBar"`
}

type GetHistoricalBarsParams struct {
	Symbols   string    `query:"symbols,required"`
	Timeframe string    `query:"timeframe,required"`
	Start     time.Time `query:"start" validate:"required"`
	End       time.Time `query:"end"`
	Feed      *string   `query:"feed,omitempty"`
	Currency  *string   `query:"currency,omitempty"`
	PageToken *string   `query:"page_token,omitempty"`
}

type GetHistoricalBarsResponse struct {
	Bars          HistoricalBarsAggregate `json:"bars"`
	NextPageToken string                  `json:"next_page_token"`
}

type HistoricalBarsAggregate map[string][]Bar

type Bar struct {
	Symbol                     string    `json:"S"`
	Open                       float64   `json:"o"`
	High                       float64   `json:"h"`
	Low                        float64   `json:"l"`
	Close                      float64   `json:"c"`
	Volume                     uint64    `json:"v"`
	VolumeWeightedAveragePrice float64   `json:"vw"`
	Timestamp                  time.Time `json:"t"`
}

type StreamStockUpdatesParams struct {
	Symbols []string `json:"symbols"`
}
