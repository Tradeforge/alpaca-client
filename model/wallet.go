package model

import "github.com/google/uuid"

type CreateFundingWalletRequest struct {
	AccountID uuid.UUID `json:"account_id"`
}

type CreateFundingWalletResponse struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

type GetFundingWalletParams struct {
	AccountID string `path:"account_id"`
}

type GetFundingWalletResponse struct {
	Status string `json:"status"`
}

type GetFundingDetailsParams struct {
	AccountID   string  `path:"account_id"`
	Currency    *string `query:"currency,omitempty"`
	PaymentType *string `query:"payment_type"`
}

type GetFundingDetailsResponse struct {
	FundingDetails []FundingDetail `json:"funding_details"`
}

type FundingDetail struct {
	AccountHolderName string `json:"account_holder_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberType string `json:"account_number_type"`
	BankAddress       string `json:"bank_address"`
	BankCountry       string `json:"bank_country"`
	BankName          string `json:"bank_name"`
	Currency          string `json:"currency"`
	PaymentType       string `json:"payment_type"`
	RoutingCode       string `json:"routing_code"`
	RoutingCodeType   string `json:"routing_code_type"`
}
