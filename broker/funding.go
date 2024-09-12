package broker

import (
	"context"
	"net/http"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
)

const (
	CreateFundingWalletPath  = "/v1beta/accounts/:account_id/funding_wallet"
	GetFundingWalletPath     = "/v1beta/accounts/:account_id/funding_wallet"
	GetFundingDetailsPath    = "/v1beta/accounts/:account_id/funding_wallet/funding_details"
	CreateInstantDepositPath = "/v1/instant_funding"
)

// FundingClient is a client for the broker account API.
type FundingClient struct {
	*client.Client
}

func (fc *FundingClient) CreateFundingWallet(ctx context.Context, data *model.CreateFundingWalletRequest, opts ...model.RequestOption) (*model.CreateFundingWalletResponse, error) {
	res := &model.CreateFundingWalletResponse{}
	err := fc.Call(ctx, http.MethodPost, CreateFundingWalletPath, nil, res, append(opts, model.Body(data))...)
	return res, err
}

func (fc *FundingClient) GetFundingWallet(ctx context.Context, params *model.GetFundingWalletParams, opts ...model.RequestOption) (*model.GetFundingWalletResponse, error) {
	res := &model.GetFundingWalletResponse{}
	err := fc.Call(ctx, http.MethodGet, GetFundingWalletPath, params, res, opts...)
	return res, err
}

func (fc *FundingClient) GetFundingDetails(ctx context.Context, params *model.GetFundingDetailsParams, opts ...model.RequestOption) (model.GetFundingDetailsResponse, error) {
	res := model.GetFundingDetailsResponse{}
	err := fc.Call(ctx, http.MethodGet, GetFundingDetailsPath, params, &res, opts...)
	return res, err
}

func (fc *FundingClient) CreateInstantFundingRequest(ctx context.Context, data *model.CreateInstantFundingRequest, opts ...model.RequestOption) (*model.CreateInstantFundingResponse, error) {
	res := &model.CreateInstantFundingResponse{}
	err := fc.Call(ctx, http.MethodPost, CreateInstantDepositPath, nil, res, append(opts, model.Body(data))...)
	return res, err
}
