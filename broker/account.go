package broker

import (
	"context"
	"net/http"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
)

const (
	CreateAccountPath          = "/v1/accounts"
	ListAccountsPath           = "/v1/accounts"
	GetAccountPath             = "/v1/accounts/:account_id"
	GetOnfidoSDKTokenPath      = "/v1/accounts/:account_id/onfido/sdk/tokens"
	UpdateOnfidoSDKOutcomePath = "/v1/accounts/:account_id/onfido/sdk"
	GetAccountHistoryPath      = "/v1/trading/accounts/:account_id/account/portfolio/history"
	GetAccountTradingDetails   = "/v1/trading/accounts/:account_id/account"
)

// FundingClient is a client for the broker account API.
type AccountClient struct {
	*client.Client
}

func (ac *AccountClient) CreateAccount(ctx context.Context, data *model.CreateAccountRequest, opts ...model.RequestOption) (*model.CreateAccountResponse, error) {
	res := &model.CreateAccountResponse{}
	err := ac.Call(ctx, http.MethodPost, CreateAccountPath, nil, res, append(opts, model.Body(data))...)
	return res, err
}

func (ac *AccountClient) ListAccounts(ctx context.Context, params model.ListAccountsParams, opts ...model.RequestOption) (model.ListAccountsResponse, error) {
	res := model.ListAccountsResponse{}
	err := ac.Call(ctx, http.MethodGet, ListAccountsPath, params, &res, opts...)
	return res, err
}

func (ac *AccountClient) GetAccount(ctx context.Context, params model.GetAccountParams, opts ...model.RequestOption) (*model.GetAccountResponse, error) {
	res := &model.GetAccountResponse{}
	err := ac.Call(ctx, http.MethodGet, GetAccountPath, params, res, opts...)
	return res, err
}

func (ac *AccountClient) GetAccountTradingDetails(ctx context.Context, params model.GetAccountTradingDetailsParams, opts ...model.RequestOption) (*model.GetAccountTradingDetailsResponse, error) {
	res := &model.GetAccountTradingDetailsResponse{}
	err := ac.Call(ctx, http.MethodGet, GetAccountTradingDetails, params, res, opts...)
	return res, err
}

func (ac *AccountClient) GetAccountHistory(ctx context.Context, params model.GetAccountHistoryParams, opts ...model.RequestOption) (*model.GetAccountHistoryResponse, error) {
	res := &model.GetAccountHistoryResponse{}
	err := ac.Call(ctx, http.MethodGet, GetAccountHistoryPath, params, res, opts...)
	return res, err
}

func (ac *AccountClient) GetOnfidoSDKToken(ctx context.Context, params model.GetOnfidoSDKTokenParams, opts ...model.RequestOption) (*model.GetOnfidoSDKTokenResponse, error) {
	res := &model.GetOnfidoSDKTokenResponse{}
	err := ac.Call(ctx, http.MethodGet, GetOnfidoSDKTokenPath, params, res, opts...)
	return res, err
}

func (ac *AccountClient) UpdateOnfidoSDKOutcome(ctx context.Context, params model.UpdateOnfidoSDKOutcomeParams, data *model.UpdateOnfidoSDKOutcomeRequest, opts ...model.RequestOption) error {
	err := ac.Call(ctx, http.MethodPatch, UpdateOnfidoSDKOutcomePath, params, http.NoBody, append(opts, model.Body(data))...)
	return err
}
