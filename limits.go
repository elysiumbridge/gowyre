package main

// LimitsRequest implements the Checkout procesing body fields from Limits API https://docs.sendwyre.com/reference#wyre-checkout
// for POST https://api.sendwyre.com/v3/widgets/limits/calculate request
type LimitsRequest struct {
	WalletType     string `json:"walletType"`
	AccountID      string `json:"accountId"`
	SourceCurrency string `json:"sourceCurrency"` // accepts USD, EUR, CAD, GBP and AUD

	Address Address `json:"address"`
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

// CalculateLimits https://docs.sendwyre.com/reference#limits-api
func (c *Client) CalculateLimits(limits *LimitsRequest) (Limits, error) {
	req, err := c.newRequest("POST", "/v3/widget/limits/calculate", limits)
	if err != nil {
		return Limits{}, err
	}
	var resp Limits
	_, err = c.do(req, &resp)
	return resp, err
}
