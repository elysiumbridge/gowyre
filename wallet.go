package gowyre

import (
	"context"
	"fmt"
)

/// Wallets are used to hold cryptocurrency funds on the Wyre platform. White label wallets are spun up on demand via the Create Wallet endpoint.

// CreateWallet body for POST https://api.sendwyre.com/v2/wallets request https://docs.sendwyre.com/reference#create-wallet
type CreateWallet struct {
	Name        string `json:"name"`
	CallbackUrl string `json:"callbackUrl"`
	Type        string `json:"type"`
	Notes       string `json:"notes"`
}

// Wallet response for POST https://api.sendwyre.com/v2/wallets request
type Wallet struct {
	Owner             string            `json:"owner"`
	Status            string            `json:"status"`
	CallbackUrl       string            `json:"callbackUrl"`
	PusherChannel     string            `json:"pusherChannel"`
	DepositAddresses  DepositAddresses  `json:"depositAddresses"`
	Balances          Balances          `json:"balances"`
	TotalBalances     TotalBalances     `json:"totalBalances"`
	AvailableBalances AvailableBalances `json:"availableBalances"`
	Notes             string            `json:"notes"`
	CreatedAt         int64             `json:"createdAt"`
	SRN               string            `json:"srn"`
	Nam               string            `json:"name"`
	ID                string            `json:"id"`
	Type              string            `json:"type"`
}

// DepositAddresses list of deposit addresses for supported currencies
// part of wallet model
type DepositAddresses struct {
	BTC  string `json:"BTC"`
	AVAX string `json:"AVAX"`
	XLM  string `json:"XLM"`
	ETH  string `json:"ETH"`
}

// Balances list of balances for supported currencies
// part of wallet model
type Balances struct {
	BTC  string `json:"BTC"`
	AVAX string `json:"AVAX"`
	XLM  string `json:"XLM"`
	ETH  string `json:"ETH"`
}

// TotalBalances list of total balances for supported currencies
// part of wallet model
type TotalBalances struct {
	BTC  string `json:"BTC"`
	AVAX string `json:"AVAX"`
	XLM  string `json:"XLM"`
	ETH  string `json:"ETH"`
}

// AvailableBalances list of available balances for supported currencies
// part of wallet model
type AvailableBalances struct {
	BTC  string `json:"BTC"`
	AVAX string `json:"AVAX"`
	XLM  string `json:"XLM"`
	ETH  string `json:"ETH"`
}

// CreateMulti body for POST https://api.sendwyre.com/v2/wallets/batch request https://docs.sendwyre.com/reference#create-multiple-wallets
type CreateMulti struct {
	Wallets []CreateWallet `json:"wallets"`
}

// UpdatePayload body for POST https://api.sendwyre.com/v2/wallet/:walletId/update request https://docs.sendwyre.com/reference#update-wallet
type UpdatePayload struct {
	WalletID    string `json:"walletId"`
	Name        string `json:"name"`
	CallbackURL string `json:"callbackUrl"`
	Notes       string `json:"notes"`
}

// ListResponse body for GET https://api.sendwyre.com/v2/wallets request
type ListResponse struct {
	Data []Wallet `json:"data"`
}

// CreateWallet https://docs.sendwyre.com/reference#create-wallet
func (c *Client) CreateWallet(ctx context.Context, wallet *CreateWallet) (Wallet, error) {
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

// LookupWallet https://docs.sendwyre.com/reference#lookup-wallet
func (c *Client) LookupWallet(ctx context.Context, walletID string) (Wallet, error) {
	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/v2/wallet/%s", walletID), nil)
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

// ListAllWallets https://docs.sendwyre.com/reference#list-all-wallets
func (c *Client) ListAllWallets(ctx context.Context) (ListResponse, error) {
	req, err := c.newRequest(ctx, "GET", "/v2/wallets", nil)
	if err != nil {
		return ListResponse{}, err
	}
	var resp ListResponse
	_, err = c.do(req, &resp)
	return resp, err
}
