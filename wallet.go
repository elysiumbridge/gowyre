package gowyre

import (
	"context"
	"errors"
	"fmt"
)

/// Wallets are used to hold cryptocurrency funds on the Wyre platform. White label wallets are spun up on demand via the Create Wallet endpoint.

var (
	ErrMissingWalletID = errors.New("No wallet identifier provided")
)

// CreateWallet body for POST https://api.sendwyre.com/v2/wallets request https://docs.sendwyre.com/reference#create-wallet
type CreateWallet struct {
	Name        string `json:"name,omitempty"`
	CallbackUrl string `json:"callbackUrl,omitempty"`
	Type        string `json:"type,omitempty"`
	Notes       string `json:"notes,omitempty"`
}

// Wallet response for POST https://api.sendwyre.com/v2/wallets request
type Wallet struct {
	Owner             string            `json:"owner,omitempty"`
	Status            string            `json:"status,omitempty"`
	CallbackUrl       string            `json:"callbackUrl,omitempty"`
	PusherChannel     string            `json:"pusherChannel,omitempty"`
	DepositAddresses  DepositAddresses  `json:"depositAddresses,omitempty"`
	Balances          Balances          `json:"balances,omitempty"`
	TotalBalances     TotalBalances     `json:"totalBalances,omitempty"`
	AvailableBalances AvailableBalances `json:"availableBalances,omitempty"`
	Notes             string            `json:"notes,omitempty"`
	CreatedAt         int64             `json:"createdAt,omitempty"`
	SRN               string            `json:"srn,omitempty"`
	Name              string            `json:"name,omitempty"`
	ID                string            `json:"id,omitempty"`
	Type              string            `json:"type,omitempty"`
}

// DepositAddresses list of deposit addresses for supported currencies
// part of wallet model
type DepositAddresses struct {
	BTC  string `json:"BTC,omitempty"`
	AVAX string `json:"AVAX,omitempty"`
	XLM  string `json:"XLM,omitempty"`
	ETH  string `json:"ETH,omitempty"`
}

// Balances list of balances for supported currencies
// part of wallet model
type Balances struct {
	BTC  float64 `json:"BTC"`
	AVAX float64 `json:"AVAX"`
	XLM  float64 `json:"XLM"`
	ETH  float64 `json:"ETH"`
}

// TotalBalances list of total balances for supported currencies
// part of wallet model
type TotalBalances struct {
	BTC  float64 `json:"BTC"`
	AVAX float64 `json:"AVAX"`
	XLM  float64 `json:"XLM"`
	ETH  float64 `json:"ETH"`
}

// AvailableBalances list of available balances for supported currencies
// part of wallet model
type AvailableBalances struct {
	BTC  float64 `json:"BTC"`
	AVAX float64 `json:"AVAX"`
	XLM  float64 `json:"XLM"`
	ETH  float64 `json:"ETH"`
}

// CreateMulti body for POST https://api.sendwyre.com/v2/wallets/batch request https://docs.sendwyre.com/reference#create-multiple-wallets
type CreateMulti struct {
	Wallets []CreateWallet `json:"wallets,omitempty"`
}

// UpdatePayload body for POST https://api.sendwyre.com/v2/wallet/:walletId/update request https://docs.sendwyre.com/reference#update-wallet
type UpdatePayload struct {
	WalletID    string `json:"walletId,omitempty"`
	Name        string `json:"name,omitempty"`
	CallbackURL string `json:"callbackUrl,omitempty"`
	Notes       string `json:"notes,omitempty"`
}

// ListResponse body for GET https://api.sendwyre.com/v2/wallets request
type ListResponse struct {
	Data []Wallet `json:"data,omitempty"`
}

// CreateWallet https://docs.sendwyre.com/reference/createwallet
func (c *Client) CreateWallet(ctx context.Context, wallet *CreateWallet) (Wallet, error) {
	if wallet.Type == "" {
		wallet.Type = "DEFAULT"
	}

	req, err := c.newRequest(ctx, "POST", "/v2/wallets", wallet)
	if err != nil {
		return Wallet{}, err
	}
	var resp Wallet
	_, err = c.do(req, &resp)
	return resp, err
}

// CreateMultipleWallet https://docs.sendwyre.com/reference#create-multiple-wallets
func (c *Client) CreateMultipleWallet(ctx context.Context, wallets *CreateMulti) (ListResponse, error) {
	req, err := c.newRequest(ctx, "POST", "/v2/wallets/batch", wallets)
	if err != nil {
		return ListResponse{}, err
	}
	var resp ListResponse
	_, err = c.do(req, &resp)
	return resp, err
}

// GetWallet https://docs.sendwyre.com/reference/getwallet
// The method receives both a walletID and name parameters. If one is specified the
// wallet will be retrieved accordingly. The walletID is used if both values are provided
func (c *Client) GetWallet(ctx context.Context, walletID, name string) (Wallet, error) {
	if walletID == "" && name == "" {
		return Wallet{}, ErrMissingWalletID
	}

	url := "/v2/wallet"
	if walletID != "" {
		url = fmt.Sprintf("%s/%s", url, walletID)
	} else {
		url = fmt.Sprintf("%s?name=%s", url, name)
	}

	req, err := c.newRequest(ctx, "GET", url, nil)
	if err != nil {
		return Wallet{}, err
	}
	var resp Wallet
	_, err = c.do(req, &resp)
	return resp, err
}

// UpdateWallet https://docs.sendwyre.com/reference#update-wallet
func (c *Client) UpdateWallet(ctx context.Context, payload *UpdatePayload) error {
	req, err := c.newRequest(ctx, "POST", fmt.Sprintf("/v2/wallet/%s/update", payload.WalletID), payload)
	if err != nil {
		return err
	}
	var resp ListResponse
	_, err = c.do(req, &resp)
	return err
}

// DeleteWallet https://docs.sendwyre.com/reference#delete-wallet
func (c *Client) DeleteWallet(ctx context.Context, walletId string) error {
	req, err := c.newRequest(ctx, "DELETE", fmt.Sprintf("/v2/wallet/%s", walletId), nil)
	if err != nil {
		return err
	}
	var resp ListResponse
	_, err = c.do(req, &resp)
	return err
}

// ListWalletsPaginated https://docs.sendwyre.com/reference#list-all-wallets
func (c *Client) ListWalletsPaginated(ctx context.Context, limit, offset int) (ListResponse, error) {
	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/v2/wallets?limit=%d&offset=%d", limit, offset), nil)
	if err != nil {
		return ListResponse{}, err
	}
	var resp ListResponse
	_, err = c.do(req, &resp)
	return resp, err
}
