package main

// CreateTransfer implements the create Transfer body fields from Transfer and Exchanges API https://docs.sendwyre.com/reference#create-transfer
// for POST https://api.sendwyre.com/v3/transfers request
type CreateTransfer struct {
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

// History response body to get transfer history https://docs.sendwyre.com/reference#transfer-history
// for GET https://api.sendwyre.com/v3/transfers request
type History struct {
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

// CreateIPValidation implements the IP validation body from Transfer and Exchanges API https://docs.sendwyre.com/reference#swap-ip-validation
// for POST https://api.sendwyre.com/v3/swaps/ipCheck request
type CreateIPValidation struct {
	IPAddress string `json:"ipAddress"`
}

// IPValidation body returned to Create IP validation
// for POST https://api.sendwyre.com/v3/swaps/ipCheck request
type IPValidation struct {
	IPAddress           string `json:"ipAddress"`
	ResolvedCountryCode string `json:"resolvedCountryCode"`
	Result              string `json:"result"`
}
