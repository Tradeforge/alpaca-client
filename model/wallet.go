package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

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

type CreateInstantFundingRequest struct {
	TargetAccountNumber string          `json:"account_no"`
	SourceAccountNumber string          `json:"source_account_no"`
	Amount              decimal.Decimal `json:"amount"`
}

type CreateInstantFundingResponse struct {
	ID               uuid.UUID         `json:"id"`
	Amount           decimal.Decimal   `json:"amount"`
	AccountNo        string            `json:"account_no"`
	SourceAccountNo  string            `json:"source_account_no"`
	TotalInterest    decimal.Decimal   `json:"total_interest"`
	RemainingPayable decimal.Decimal   `json:"remaining_payable"`
	Interests        []DepositInterest `json:"interests"`
	Fees             []DepositFee      `json:"fees"`
	SystemDate       string            `json:"system_date"`
	Deadline         string            `json:"deadline"`
	Status           string            `json:"status"`
	CreatedAt        time.Time         `json:"created_at"`
}

type DepositInterest struct {
	ID           uuid.UUID       `json:"id"`
	Date         string          `json:"date"`
	Amount       decimal.Decimal `json:"amount"`
	Status       string          `json:"status"`
	CreatedAt    time.Time       `json:"created_at"`
	ReconciledAt time.Time       `json:"reconciled_at"`
}

type DepositFee struct {
	ID     uuid.UUID       `json:"id"`
	Amount decimal.Decimal `json:"amount"`
	Type   string          `json:"type"`
}
