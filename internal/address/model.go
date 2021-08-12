package address

// Model customer address related card info
type Model struct {
	Street1 string `json:"street1"`
	City    string `json:"city"`
	// State state code
	State string `json:"state"`
	// PostalCode only numbers
	PostalCode string `json:"postalCode"`
	// Country alpha2 country code
	Country string `json:"country"`
}
