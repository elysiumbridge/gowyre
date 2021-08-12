package main

// Model implements the Checkout procesing body fields from Wyre Checkout API https://docs.sendwyre.com/reference#wyre-checkout
// for POST https://api.sendwyre.com/v3/orders/reserve request
type Model struct {
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

// LimitsRequest implements the Checkout procesing body fields from Limits API https://docs.sendwyre.com/reference#wyre-checkout
// for POST https://api.sendwyre.com/v3/widgets/limits/calculate request
type LimitsRequest struct {
	WalletType     string `json:"walletType"`
	AccountID      string `json:"accountId"`
	SourceCurrency string `json:"sourceCurrency"` // accepts USD, EUR, CAD, GBP and AUD

	Address Address `json:"address"`
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

// Limits body for limits request listing user limits
type Limits struct {
	DailyTotalSpent  float64 `json:"dailyTotalSpent"`
	WeeklyTotalSpent float64 `json:"weeklyTotalSpent"`
	YearlyTotalSpent float64 `json:"yearlyTotalSpent"`
	DailyTotal       float64 `json:"dailyTotal"`
	WeeklyTotal      float64 `json:"weeklyTotal"`
	YearlyTotal      float64 `json:"yearlyTotal"`
	DailyRemaining   float64 `json:"dailyRemaining"`
	WeeklyRemaining  float64 `json:"weeklyRemaining"`
	YearlyRemaining  float64 `json:"yearlyRemaining"`
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
