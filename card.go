package main

// Card implements the Card procesing body fields from Card Processing API https://docs.sendwyre.com/reference#white-label-card-processing-api
// for POST https://api.sendwyre.com/v3/debitcard/process/partner request
type Card struct {
	DebitCard DebitCard `json:"debitCard"`

	ReservationID     string `json:"reservationId"`
	Amount            string `json:"amount"`
	SourceCurrency    string `json:"sourceCurrency"`
	DestCurrency      string `json:"destCurrency"`
	Dest              string `json:"dest"`
	ReferrerAccountId string `json:"referrerAccountId"`
	GivenName         string `json:"givenName"`
	FamilyName        string `json:"familyName"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	ReferenceId       string `json:"referenceId"`

	Address Address `json:"address"`
}

// ProcessResponse is the response from https://api.sendwyre.com/v3/debitcard/process/partner POST request
type ProcessResponse struct {
	ID                string `json:"id"`
	CreatedAt         string `json:"createdAt"`
	Owner             string `json:"owner"`
	Status            string `json:"status"`
	TransferID        string `json:"transferId"`
	SourceAmount      string `json:"sourceAmount"`
	SourceCurrency    string `json:"sourceCurrency"`
	DestCurrency      string `json:"destCurrency"`
	Dest              string `json:"dest"`
	WalletType        string `json:"walletType"`
	Email             string `json:"email"`
	ErrorMessage      string `json:"errorMessage"`
	AccountId         string `json:"accountId"`
	PaymentMethodName string `json:"paymentMethodName"`
}

// DebitCard debit card details from Card Processing API
type DebitCard struct {
	Number string `json:"number"`
	Year   string `json:"year"`
	Month  string `json:"month"`
	CVV    string `json:"cvv"`
}

// Authorize body for POST https://api.sendwyre.com/v3/debitcard/authorize/partner request
// that returns if card is authorized https://docs.sendwyre.com/reference#authorize-card
type Authorize struct {
	// Type "SMS", "CARD2FA" or "ALL"
	Type          string `json:"type"`
	WalletOrderId string `json:"walletOrderId"`
	Reservation   string `json:"reservation"`
	// SMS A 6 digits sent to the user's mobile device.
	SMS string `json:"sms"`
	// Card2FA A 6 digit code that's added immediately to the end user's bank account.
	Card2FA string `json:"card2fa"`
}

// Authorize body response to POST https://api.sendwyre.com/v3/debitcard/authorize/partner request
// that returns authorization response
type AuthorizeResponse struct {
	WalletOrderId string `json:"walletOrderId"`
	Success       string `json:"success"`
}

// Authorized body response to GET https://api.sendwyre.com/v3/debitcard/authorization/orderId request
// that returns if card is authorized https://docs.sendwyre.com/reference#authorize-card
type Authorized struct {
	WalletOrderId       string `json:"walletOrderId"`
	SmsNeeded           string `json:"smsNeeded"`
	Card2faNeeded       string `json:"card2faNeeded"`
	Authorization3dsURL string `json:"authorization3dsUrl"`
}

// Refund body response to POST https://api.sendwyre.com/v3/orders/:orderId/refund/partner request
// that returns refund response
type RefundResponse struct {
	WalletOrderId string `json:"walletOrderId"`
	CreatedAt     string `json:"createdAt"`
	Requester     string `json:"requester"`
	Status        string `json:"status"`
	FinishedAt    string `json:"finishedAt"`
	FailedReason  string `json:"failedReason"`
}
