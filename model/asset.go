package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ListAssetsParams struct {
	// AssetClass is the asset class to filter by. Valid values are "us_equity", "crypto"
	AssetClass string `query:"asset_class"`
	// AssetStatus is the status of the asset. Valid values are "active", "inactive", or "all".
	Status string `json:"status"`
}

type ListAssetsResponse []Asset

type Asset struct {
	ID                     uuid.UUID       `json:"id"`
	Class                  string          `json:"class"`
	Exchange               string          `json:"exchange"`
	Symbol                 string          `json:"symbol"`
	Name                   string          `json:"name"`
	Status                 string          `json:"status"`
	Tradable               bool            `json:"tradable"`
	Marginable             bool            `json:"marginable"`
	Shortable              bool            `json:"shortable"`
	EasyToBorrow           bool            `json:"easy_to_borrow"`
	Fractionable           bool            `json:"fractionable"`
	MinOrderSize           decimal.Decimal `json:"min_order_size"`
	MinTradeIncrement      decimal.Decimal `json:"min_trade_increment"`
	PriceIncrement         decimal.Decimal `json:"price_increment"`
	MarginRequirementLong  decimal.Decimal `json:"margin_requirement_long"`
	MarginRequirementShort decimal.Decimal `json:"margin_requirement_short"`
	Attributes             []string        `json:"attributes"`
}
