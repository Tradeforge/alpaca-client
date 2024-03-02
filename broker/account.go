package broker

import (
	"context"
	"net/http"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
)

const (
	CreateAccountPath     = "/v1/accounts"
	ListAccountsPath      = "/v1/accounts"
	GetAccountPath        = "/v1/accounts/:id"
	GetAccountHistoryPath = "/v1/trading/accounts/:id/account/portfolio/history"
)

// AccountClient is a client for the broker account API.
type AccountClient struct {
	client.Client
}

func (ac *AccountClient) CreateAccount(ctx context.Context, data *model.CreateAccountRequest, opts ...model.RequestOption) (*model.CreateAccountResponse, error) {
	res := &model.CreateAccountResponse{}
	err := ac.Call(ctx, http.MethodPost, CreateAccountPath, nil, res, model.Body(data))
	return res, err
}

func (ac *AccountClient) ListAccounts(ctx context.Context, params *model.ListAccountsParams, opts ...model.RequestOption) (model.ListAccountsResponse, error) {
	res := model.ListAccountsResponse{}
	err := ac.Call(ctx, http.MethodGet, ListAccountsPath, params, &res)
	return res, err
}

func (ac *AccountClient) GetAccount(ctx context.Context, params *model.GetAccountParams, opts ...model.RequestOption) (*model.GetAccountResponse, error) {
	res := &model.GetAccountResponse{}
	err := ac.Call(ctx, http.MethodGet, GetAccountPath, params, res)
	return res, err
}

func (ac *AccountClient) GetAccountHistory(ctx context.Context, params *model.GetAccountHistoryParams, opts ...model.RequestOption) (*model.GetAccountHistoryResponse, error) {
	res := &model.GetAccountHistoryResponse{}
	err := ac.Call(ctx, http.MethodGet, GetAccountHistoryPath, params, res, opts...)
	return res, err
}
