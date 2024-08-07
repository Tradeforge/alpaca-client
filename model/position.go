package model

import "github.com/shopspring/decimal"

type GetOpenPositionBySymbolParams struct {
	AccountID string `path:"account_id"`
	Symbol    string `path:"symbol"`
}

type GetOpenPositionResponse struct {
	AssetID                string          `json:"asset_id"`
	Symbol                 string          `json:"symbol"`
	Exchange               string          `json:"exchange"`
	AssetClass             string          `json:"asset_class"`
	AssetMarginable        bool            `json:"asset_marginable"`
	AverageEntryPrice      decimal.Decimal `json:"avg_entry_price"`
	Quantity               decimal.Decimal `json:"qty"`
	Side                   string          `json:"side"`
	MarketValue            decimal.Decimal `json:"market_value"`
	CostBasis              decimal.Decimal `json:"cost_basis"`
	UnrealizedPL           decimal.Decimal `json:"unrealized_pl"`
	UnrealizedPLPC         decimal.Decimal `json:"unrealized_plpc"`
	UnrealizedIntradayPL   decimal.Decimal `json:"unrealized_intraday_pl"`
	UnrealizedIntradayPLPC decimal.Decimal `json:"unrealized_intraday_plpc"`
	CurrentPrice           decimal.Decimal `json:"current_price"`
	LastDayPrice           decimal.Decimal `json:"lastday_price"`
	ChangeToday            decimal.Decimal `json:"change_today"`
	USD                    struct {
		AverageEntryPrice      decimal.Decimal `json:"avg_entry_price"`
		MarketValue            decimal.Decimal `json:"market_value"`
		CostBasis              decimal.Decimal `json:"cost_basis"`
		UnrealizedPL           decimal.Decimal `json:"unrealized_pl"`
		UnrealizedPLPC         decimal.Decimal `json:"unrealized_plpc"`
		UnrealizedIntradayPL   decimal.Decimal `json:"unrealized_intraday_pl"`
		UnrealizedIntradayPLPC decimal.Decimal `json:"unrealized_intraday_plpc"`
		CurrentPrice           decimal.Decimal `json:"current_price"`
		LastDayPrice           decimal.Decimal `json:"lastday_price"`
		ChangeToday            decimal.Decimal `json:"change_today"`
	} `json:"usd"`
}
