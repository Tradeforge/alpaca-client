package model

type Transfer struct {
	ID     string         `json:"id"`
	Status TransferStatus `json:"status"`
}

// TransferStatus defines the status of a transfer.
// See https://github.com/alpacahq/alpaca-docs/blob/master/content/api-references/broker-api/funding/transfers.md#enumtransferstatus.
type TransferStatus string
