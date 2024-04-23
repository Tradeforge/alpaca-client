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
	AskSize   int64     `json:"as"`
	BidPrice  float64   `json:"bp"`
	BidSize   int64     `json:"bs"`
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
	Size      int64     `json:"s"`
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

/*
Example Response:

	{
	  "bars": {
	    "AAPL": [
	      {
	        "t": "2022-01-03T09:00:00Z",
	        "o": 178.26,
	        "h": 178.26,
	        "l": 178.21,
	        "c": 178.21,
	        "v": 1118,
	        "n": 65,
	        "vw": 178.235733
	      }
	    ]
	  },
	  "next_page_token": "QUFQTHxNfDIwMjItMDEtMDNUMDk6MDA6MDAuMDAwMDAwMDAwWg=="
	}
*/
type GetHistoricalBarsResponse struct {
	Bars          HistoricalBarsAggregate `json:"bars"`
	NextPageToken string                  `json:"next_page_token"`
}

type HistoricalBarsAggregate map[string][]Bar

type Bar struct {
	Open           float64   `json:"o"`
	High           float64   `json:"h"`
	Low            float64   `json:"l"`
	Close          float64   `json:"c"`
	Volume         int64     `json:"v"`
	VolumeWeighted float64   `json:"vw"`
	Timestamp      time.Time `json:"t"`
}
