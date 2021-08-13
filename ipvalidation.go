package gowyre

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

// SwapIPValidation https://docs.sendwyre.com/reference#swap-ip-validation
func (c *Client) SwapIPValidation(IPAddress string) (IPValidation, error) {
	req, err := c.newRequest("POST", "/v3/swaps/ipCheck", &CreateIPValidation{IPAddress})
	if err != nil {
		return IPValidation{}, err
	}
	var resp IPValidation
	_, err = c.do(req, &resp)
	return resp, err
}
