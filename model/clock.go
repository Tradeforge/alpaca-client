package model

import "time"

type GetMarketClockResponse struct {
	Timestamp time.Time `json:"timestamp"`
	IsOpen    bool      `json:"is_ope"`
	NextOpen  time.Time `json:"next_open"`
	NextClose time.Time `json:"next_close"`
}
