package gowyre

import (
	"context"
	"fmt"
)

// CreateTransfer implements the create Transfer body fields from Transfer and Exchanges API https://docs.sendwyre.com/reference#create-transfer
// for POST https://api.sendwyre.com/v3/transfers request
type CreateTransfer struct {
	CustomId       string  `json:"customId"`
	Source         string  `json:"source"`
	SourceCurrency string  `json:"sourceCurrency"`
	SourceAmount   float64 `json:"sourceAmount"`
	Dest           string  `json:"dest"`
	DestCurrency   string  `json:"destCurrency"`
	Message        string  `json:"message"`
	AutoConfirm    bool    `json:"autoConfirm"`
}

// Model transfer body returned to Create new transfer
// for POST https://api.sendwyre.com/v3/transfers request
type Transfer struct {
	Owner              string          `json:"owner"`
	Source             string          `json:"source"`
	SourceAmount       float64         `json:"sourceAmount"`
	SourceCurrency     string          `json:"sourceCurrency"`
	DestCurrency       string          `json:"destCurrency"`
	CustomID           *string         `json:"customId"`
	CompletedAt        *uint64         `json:"completedAt"`
	CancelledAt        *uint64         `json:"cancelledAt"`
	FailureReason      *string         `json:"failureReason"`
	CreatedAt          uint64          `json:"createdAt"`
	ExpiresAt          uint64          `json:"expiresAt"`
	UpdatedAt          *uint64         `json:"updatedAt"`
	Status             string          `json:"status"`
	ExchangeRate       *float64        `json:"exchangeRate"`
	DestAmount         float64         `json:"destAmount"`
	Fees               Fee             `json:"fees"`
	TotalFees          float64         `json:"totalFees"`
	BlockchainTx       Tx              `json:"blockchainTx"`
	ReversalReason     *string         `json:"reversalReason"`
	ReversingSubStatus *string         `json:"reversingSubStatus"`
	Dest               string          `json:"dest"`
	StatusHistories    []StatusHistory `json:"statusHistories"`
	PendingSubStatus   string          `json:"pendingSubStatus"`
	Message            string          `json:"message"`
	ID                 string          `json:"id"`
}

// StatusHistory part of transfer model
type StatusHistory struct {
	ID           string  `json:"id"`
	TransferID   string  `json:"transferId"`
	CreatedAt    uint64  `json:"createdAt"`
	Type         string  `json:"type"`
	StatusOrder  int64   `json:"statusOrder"`
	StatusDetail string  `json:"statusDetail"`
	State        string  `json:"state"`
	FailedState  *string `json:"failedState"`
}

// TransfersHistory response body to get transfer history https://docs.sendwyre.com/reference#transfer-history
// for GET https://api.sendwyre.com/v3/transfers request
type TransfersHistory struct {
	Data []struct {
		ClosedAt       uint64   `json:"closedAt"`
		CreatedAt      uint64   `json:"createdAt"`
		ID             string   `json:"id"`
		CustomID       *string  `json:"customId"`
		Source         string   `json:"source"`
		Dest           string   `json:"dest"`
		SourceCurrency string   `json:"sourceCurrency"`
		DestCurrency   string   `json:"destCurrency"`
		SourceAmount   string   `json:"sourceAmount"`
		DestAmount     float64  `json:"destAmount"`
		Fees           Fee      `json:"fees"`
		SourceName     string   `json:"sourceName"`
		DestName       string   `json:"destName"`
		Status         string   `json:"status"`
		Message        *string  `json:"message"`
		ExchangeRate   *float64 `json:"exchangeRate"`
		BlockchainTx   Tx       `json:"blockchainTx"`
		DestNickname   string   `json:"destNickname"`
	} `json:"data"`
	Position        int64 `json:"position"`
	RecordsTotal    int64 `json:"recordsTotal"`
	RecordsFiltered int64 `json:"recordsFiltered"`
}

// CreateTransfer https://docs.sendwyre.com/reference#create-transfer
func (c *Client) CreateTransfer(ctx context.Context, transfer *CreateTransfer) (Transfer, error) {
	req, err := c.newRequest(ctx, "POST", "/v3/transfers", transfer)
	if err != nil {
		return Transfer{}, err
	}
	var resp Transfer
	_, err = c.do(req, &resp)
	return resp, err
}

// ConfirmTransfer https://docs.sendwyre.com/reference#confirm-transfer
func (c *Client) ConfirmTransfer(ctx context.Context, transferID string) (Transfer, error) {
	req, err := c.newRequest(ctx, "POST", fmt.Sprintf("/v3/transfers/%s/confirm", transferID), nil)
	if err != nil {
		return Transfer{}, err
	}
	var resp Transfer
	_, err = c.do(req, &resp)
	return resp, err
}

// GetTransfer https://docs.sendwyre.com/reference#get-transfer
func (c *Client) GetTransfer(ctx context.Context, transferID string) (Transfer, error) {
	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/v3/transfers/%s", transferID), nil)
	if err != nil {
		return Transfer{}, err
	}
	var resp Transfer
	_, err = c.do(req, &resp)
	return resp, err
}

// GetTransferByCustomID https://docs.sendwyre.com/reference#get-transfer-by-custom-id
func (c *Client) GetTransferByCustomID(ctx context.Context, customID string) (Transfer, error) {
	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/v2/transfers?customId=%s", customID), nil)
	if err != nil {
		return Transfer{}, err
	}
	var resp Transfer
	_, err = c.do(req, &resp)
	return resp, err
}

// GetTransfersHistory https://docs.sendwyre.com/reference#transfer-history
func (c *Client) GetTransfersHistory(ctx context.Context) (TransfersHistory, error) {
	req, err := c.newRequest(ctx, "GET", "/v3/transfers", nil)
	if err != nil {
		return TransfersHistory{}, err
	}
	var resp TransfersHistory
	_, err = c.do(req, &resp)
	return resp, err
}
