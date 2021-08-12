package main

import "fmt"

// Checkout implements the Checkout procesing body fields from Wyre Checkout API https://docs.sendwyre.com/reference#wyre-checkout
// for POST https://api.sendwyre.com/v3/orders/reserve request
type Checkout struct {
	PaymentMethod      string   `json:"paymentMethod"`
	Amount             string   `json:"amount"`
	SourceCurrency     string   `json:"sourceCurrency"`
	RedirectUrl        string   `json:"redirectUrl"`
	FailureRedirectUrl string   `json:"failureRedirectUrl"`
	ReferrerAccountId  string   `json:"referrerAccountId"`
	Country            string   `json:"country"`
	LockFields         []string `json:"lockFields"`
}

// ReservationResponse returns response to reservation request
type ReservationResponse struct {
	URL         string `json:"url"`
	Reservation string `json:"reservation"`
}

// QuotationRequest implements the Quotation procesing body fields from Wyre Checkout API https://docs.sendwyre.com/reference#wallet-order-quotation
// for POST https://api.sendwyre.com/v3/orders/quote/partner request
type QuotationRequest struct {
	Amount         string `json:"amount"`
	SourceCurrency string `json:"sourceCurrency"`
	DestCurrency   string `json:"destCurrency"`
	Dest           string `json:"dest"`
	AccountId      string `json:"accountId"`
	Country        string `json:"country"`
}

// Quotation model list the quotation changes from base currency
type Quotation struct {
	SourceCurrency string        `json:"sourceCurrency"`
	SourceAmount   float64       `json:"sourceAmount"`
	DestCurrency   string        `json:"destCurrency"`
	DestAmount     float64       `json:"destAmount"`
	ExchangeRate   float64       `json:"exchangeRate"`
	Equivalencies  Equivalencies `json:"equivalencies"`
	Fees           Fee           `json:"fees"`
}

// Equivalencies list of currency quotation pairs
type Equivalencies struct {
	CAD  float64 `json:"CAD"`
	USDC float64 `json:"USDC"`
	BTC  float64 `json:"BTC"`
	ETH  float64 `json:"ETH"`
	GBP  float64 `json:"GBP"`
	DAI  float64 `json:"DAI"`
	AUD  float64 `json:"AUD"`
	EUR  float64 `json:"EUR"`
	WETH float64 `json:"WETH"`
	USD  float64 `json:"USD"`
	MXN  float64 `json:"MXN"`
}

// OrderDetail body response to GET https://api.sendwyre.com/v3/orders/orderId request
// Reference https://docs.sendwyre.com/reference#wallet-order-details
type OrderDetail struct {
	ID             string  `json:"id"`
	CreatedAt      uint64  `json:"createdAt"`
	Owner          string  `json:"owner"`
	Status         string  `json:"status"`
	SourceAmount   float64 `json:"sourceAmount"`
	SourceCurrency string  `json:"sourceCurrency"`
	DestCurrency   string  `json:"destCurrency"`
	Dest           string  `json:"dest"`
	WalletType     string  `json:"walletType"`
	TransferID     *string `json:"transferId"`
	ErrorMessage   *string `json:"errorMessage"`
	AccountID      string  `json:"accountId"`
}

// OrderFullDetail body response to GET https://api.sendwyre.com/v3/orders/id/full request
// Reference https://docs.sendwyre.com/reference#wallet-order-full
type OrderFullDetail struct {
	ID             string  `json:"id"`
	CreatedAt      uint64  `json:"createdAt"`
	Owner          string  `json:"owner"`
	Status         string  `json:"status"`
	SourceAmount   float64 `json:"sourceAmount"`
	SourceCurrency string  `json:"sourceCurrency"`
	DestCurrency   string  `json:"destCurrency"`
	Dest           string  `json:"dest"`
	WalletType     string  `json:"walletType"`
	TransferID     *string `json:"transferId"`
	ErrorMessage   *string `json:"errorMessage"`
	AccountID      string  `json:"accountId"`

	OrderType               string `json:"orderType"`
	PurchaseAmount          string `json:"purchaseAmount"`
	AuthCodesRequested      bool   `json:"authCodesRequested"`
	ErrorCategory           string `json:"errorCategory"`
	ErrorCode               string `json:"errorCode"`
	FailureReason           string `json:"failureReason"`
	PaymentNetworkErrorCode string `json:"paymentNetworkErrorCode"`
	InternalErrorCode       string `json:"internalErrorCode"`
}

// OrdersList body response to GET https://api.sendwyre.com/v3/orders/list request
// Reference https://docs.sendwyre.com/reference#list-orders-paged
type OrdersList struct {
	Data []OrderFullDetail `json:"data"`
}

// TrackWidgetOrder body response to GET https://api.sendwyre.com/v2/transfer/transferId/track request
// Reference https://docs.sendwyre.com/reference#track-wallet-order
type TrackWidgetOrder struct {
	TransferId               string          `json:"transferId"`
	FeeCurrency              string          `json:"feeCurrency"`
	Fee                      float64         `json:"fee"`
	Fees                     Fee             `json:"fees"`
	SourceCurrency           string          `json:"sourceCurrency"`
	DestCurrency             string          `json:"destCurrency"`
	SourceAmount             float64         `json:"sourceAmount"`
	DestAmount               float64         `json:"destAmount"`
	DestSrn                  string          `json:"destSrn"`
	From                     string          `json:"from"`
	To                       *string         `json:"to"`
	Rate                     float64         `json:"rate"`
	CustomID                 *string         `json:"customId"`
	Status                   string          `json:"status"`
	BlockchainNetworkTx      *string         `json:"blockchainNetworkTx"`
	Message                  *string         `json:"message"`
	TransferHistoryEntryType string          `json:"transferHistoryEntryType"`
	SuccessTimeline          []StateTimeline `json:"successTimeline"`
	FailedTimeline           []StateTimeline `json:"failedTimeline"`
	FailureReason            *string         `json:"failureReason"`
	ReversalReason           *string         `json:"reversalReason"`
}

// StateTimeline part of TrackWidgetOrder defines timeline statuses
type StateTimeline struct {
	StatusDetails string `json:"statusDetails"`
	State         string `json:"state"`
	CreatedAt     int64  `json:"createdAt"`
}

// SupportedCountries body response to GET https://api.sendwyre.com/v3/widget/supportedCountries request
// Reference https://docs.sendwyre.com/reference#country-list
type SupportedCountries []SupportedCountry

// SupportedCountry element in SupportedCountries response to GET https://api.sendwyre.com/v3/widget/supportedCountries request
type SupportedCountry struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	PhonePrefix string `json:"phonePrefix"`
	StateLabel  string `json:"stateLabel"`
}

// RateLockedReservation body response to GET https://api.sendwyre.com/v3/orders/reservation/reservationId request
// Reference https://docs.sendwyre.com/reference#rate-locked-reservation
type RateLockedReservation struct {
	Amount            *string  `json:"amount"`
	SourceCurrency    string   `json:"sourceCurrency"`
	DestCurrency      string   `json:"destCurrency"`
	Dest              string   `json:"dest"`
	ReferrerAccountId string   `json:"referrerAccountId"`
	SourceAmount      *float64 `json:"sourceAmount"`
	DestAmount        float64  `json:"destAmount"`
	AmountIncludeFees bool     `json:"amountIncludeFees"`

	Address

	FirstName          string   `json:"firstName"`
	LastName           string   `json:"lastName"`
	Phone              string   `json:"phone"`
	Email              string   `json:"email"`
	LockFields         []string `json:"lockFields"`
	RedirectUrl        string   `json:"redirectUrl"`
	FailureRedirectUrl string   `json:"failureRedirectUrl"`
	PaymentMethod      string   `json:"paymentMethod"`
	ReferenceId        *string  `json:"referenceId"`
	HideTrackBtn       *string  `json:"hideTrackBtn"`

	Quote            Quote `json:"quote"`
	QuoteLockRequest bool  `json:"quoteLockRequest"`
}

// Quote part of RateLockedReservation
type Quote struct {
	SourceCurrency          string        `json:"sourceCurrency"`
	SourceAmount            float64       `json:"sourceAmount"`
	SourceAmountWithoutFees float64       `json:"sourceAmountWithoutFees"`
	DestCurrency            string        `json:"destCurrency"`
	DestAmount              float64       `json:"destAmount"`
	ExchangeRate            float64       `json:"exchangeRate"`
	Equivalencies           Equivalencies `json:"equivalencies"`
	Fees                    Fee           `json:"fees"`
}

// CreateWalletOrderReservation https://docs.sendwyre.com/reference#wallet-order-reservations
func (c *Client) CreateWalletOrderReservation(payload *Checkout) (ReservationResponse, error) {
	req, err := c.newRequest("POST", "/v3/orders/reserve", payload)
	if err != nil {
		return ReservationResponse{}, err
	}
	var resp ReservationResponse
	_, err = c.do(req, &resp)
	return resp, err
}

// CreateWalletOrderQuotation https://docs.sendwyre.com/reference#wallet-order-details
func (c *Client) CreateWalletOrderQuotation(payload *QuotationRequest) (Quotation, error) {
	req, err := c.newRequest("POST", "/v3/orders/quote/partner", payload)
	if err != nil {
		return Quotation{}, err
	}
	var resp Quotation
	_, err = c.do(req, &resp)
	return resp, err
}

// GetWalletOrderFull https://docs.sendwyre.com/reference#wallet-order-full
func (c *Client) GetWalletOrderFull(orderId string) (OrderFullDetail, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/v3/orders/%s/full", orderId), nil)
	if err != nil {
		return OrderFullDetail{}, err
	}
	var resp OrderFullDetail
	_, err = c.do(req, &resp)
	return resp, err
}

// GetWalletOrdersList https://docs.sendwyre.com/reference#list-orders-paged
func (c *Client) GetWalletOrdersList() (OrdersList, error) {
	req, err := c.newRequest("GET", "/v3/orders/list", nil)
	if err != nil {
		return OrdersList{}, err
	}
	var resp OrdersList
	_, err = c.do(req, &resp)
	return resp, err
}

// GetTrackWidgetOrder https://docs.sendwyre.com/reference#track-wallet-order
func (c *Client) GetTrackWidgetOrder(transferId string) (TrackWidgetOrder, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/v2/transfer/%s/track", transferId), nil)
	if err != nil {
		return TrackWidgetOrder{}, err
	}
	var resp TrackWidgetOrder
	_, err = c.do(req, &resp)
	return resp, err
}

// GetSupportedCountries https://docs.sendwyre.com/reference#list-orders-paged
func (c *Client) GetSupportedCountries() (SupportedCountries, error) {
	req, err := c.newRequest("GET", "/v3/widget/supportedCountries", nil)
	if err != nil {
		return SupportedCountries{}, err
	}
	var resp SupportedCountries
	_, err = c.do(req, &resp)
	return resp, err
}

// GetRateLockedReservation https://docs.sendwyre.com/reference#rate-locked-reservation
func (c *Client) GetRateLockedReservation(reservationId string) (RateLockedReservation, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/v3/orders/reservation/%s", reservationId), nil)
	if err != nil {
		return RateLockedReservation{}, err
	}
	var resp RateLockedReservation
	_, err = c.do(req, &resp)
	return resp, err
}
