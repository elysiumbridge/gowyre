package gowyre

import (
	"errors"
	"fmt"
)

const (
	DefaultRate = "DIVISOR"
)

var ErrUnsupportedRate = errors.New("the provided exchange rate type is not supported")

// GetExchangeRates https://docs.sendwyre.com/reference#live-exchange-rates
func (c *Client) GetExchangeRates(as string) (interface{}, error) {
	if as == "" {
		as = DefaultRate
	}
	if as != "DIVISOR" && as != "MULTIPLIER" && as != "PRICED" {
		return nil, ErrUnsupportedRate
	}

	req, err := c.newRequest("GET", fmt.Sprintf("/v3/rates?as=%s", as), nil)
	if err != nil {
		return nil, err
	}
	var resp interface{}
	_, err = c.do(req, &resp)
	return resp, err
}
