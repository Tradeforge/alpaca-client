package model

import (
	"time"
)

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
	Bars          map[string][]HistoricalBar `json:"bars"`
	NextPageToken string                     `json:"next_page_token"`
}

type HistoricalBar struct {
	Open           float64 `json:"o"`
	High           float64 `json:"h"`
	Low            float64 `json:"l"`
	Close          float64 `json:"c"`
	Volume         int64   `json:"v"`
	VolumeWeighted float64 `json:"vw"`
	// Timestamp in RFC-3339 format with nanosecond precision.
	Timestamp time.Time `json:"t"`
}
