package broker

import (
	"context"
	"net/http"

	"go.tradeforge.dev/alpaca/client"
	"go.tradeforge.dev/alpaca/model"
)

const (
	UploadDocumentPath = "/v1/accounts/:account_id/documents/upload"
)

// DocumentClient is a client for the document API.
type DocumentClient struct {
	*client.Client
}

func (ac *AccountClient) UploadDocument(ctx context.Context, data *model.UploadDocumentRequest, params *model.UploadDocumentParams, opts ...model.RequestOption) error {
	res := &model.UploadDocumentResponse{}
	err := ac.Call(ctx, http.MethodPost, UploadDocumentPath, params, res, model.Body(data))
	return err
}
