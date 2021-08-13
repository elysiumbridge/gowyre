package gowyre

import "fmt"

// CreateSwap implements the create Swap body fields from Transfer and Exchanges API https://docs.sendwyre.com/reference#create-swap
// for POST https://api.sendwyre.com/v3/swaps request
type CreateSwap struct {
	Dest           string `json:"dest"`
	SourceCurrency string `json:"sourceCurrency"`
	DestCurrency   string `json:"destCurrency"`
	RefundTo       string `json:"refundTo"`
	NotifyUrl      string `json:"notifyUrl"`
	IPAddress      string `json:"ipAddress"`
}

// Swap body returned to Create new swap
// for POST https://api.sendwyre.com/v3/swaps request
type Swap struct {
	ID             string         `json:"id"`
	Owner          string         `json:"owner"`
	Status         string         `json:"status"`
	SourceCurrency string         `json:"sourceCurrency"`
	DestCurrency   string         `json:"destCurrency"`
	Dest           string         `json:"dest"`
	RefundTo       string         `json:"refundTo"`
	FundingAddress string         `json:"fundingAddress"`
	CreatedAt      uint64         `json:"createdAt"`
	UpdatedAt      uint64         `json:"updatedAt"`
	ExpiresAt      uint64         `json:"expiresAt"`
	SwapTransfers  []SwapTransfer `json:"swapTransfers"`
}

// SwapTransfer tranfer part of the swap
type SwapTransfer struct {
	ID             string  `json:"id"`
	SwapID         string  `json:"swapId"`
	TransferID     string  `json:"transferId"`
	Type           string  `json:"type"`
	Reason         *string `json:"reason"`
	RefundedTo     *string `json:"refundedTo"`
	CreatedAt      uint64  `json:"createdAt"`
	Source         string  `json:"source"`
	Dest           string  `json:"dest"`
	SourceCurrency string  `json:"sourceCurrency"`
	DestCurrency   string  `json:"destCurrency"`
	SourceAmount   float64 `json:"sourceAmount"`
	DestAmount     float64 `json:"destAmount"`
	TransferStatus string  `json:"transferStatus"`
	NetworkTxId    *string `json:"networkTxId"`
}

// CreateSwap https://docs.sendwyre.com/reference#create-swap
func (c *Client) CreateSwap(swap *CreateSwap) (Swap, error) {
	req, err := c.newRequest("POST", "/v3/swaps", swap)
	if err != nil {
		return Swap{}, err
	}
	var resp Swap
	_, err = c.do(req, &resp)
	return resp, err
}

// GetSwap https://docs.sendwyre.com/reference#get-swap
func (c *Client) GetSwap(swapID string) (Swap, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/v3/swaps/%s", swapID), nil)
	if err != nil {
		return Swap{}, err
	}
	var resp Swap
	_, err = c.do(req, &resp)
	return resp, err
}
