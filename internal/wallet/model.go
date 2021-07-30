package wallet

/// Wallets are used to hold cryptocurrency funds on the Wyre platform. White label wallets are spun up on demand via the Create Wallet endpoint.

// Create body for POST https://api.sendwyre.com/v2/wallets request https://docs.sendwyre.com/reference#create-wallet
type Create struct {
	Name        string `json:"name"`
	CallbackUrl string `json:"callbackUrl"`
	Type        string `json:"type"`
	Notes       string `json:"notes"`
}

// Model response for POST https://api.sendwyre.com/v2/wallets request
type Model struct {
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
	Wallets []Create `json:"wallets"`
}

// Update body for POST https://api.sendwyre.com/v2/wallet/:walletId/update request https://docs.sendwyre.com/reference#update-wallet
type Update struct {
	WalletID    string `json:"walletId"`
	Name        string `json:"name"`
	CallbackURL string `json:"callbackUrl"`
	Notes       string `json:"notes"`
}

// ListResponse body for GET https://api.sendwyre.com/v2/wallets request
type ListResponse struct {
	Data []Model `json:"data"`
}
