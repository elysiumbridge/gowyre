package gowyre

import (
	"context"
	"errors"
	"fmt"
)

type Rate string

const (
	Divisor    Rate = "DIVISOR"
	Multiplier Rate = "MULTIPLIER"
	Priced     Rate = "PRICED"

	DefaultRate = Divisor
)

var ErrUnsupportedRate = errors.New("the provided exchange rate type is not supported")

// GetExchangeRates https://docs.sendwyre.com/reference#live-exchange-rates
func (c *Client) GetExchangeRates(ctx context.Context, as Rate) (map[string]float64, error) {
	if as == "" {
		as = DefaultRate
	}
	if as != Divisor && as != Multiplier && as != Priced {
		return nil, ErrUnsupportedRate
	}

	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/v3/rates?as=%s", as), nil)
	if err != nil {
		return nil, err
	}
	var resp map[string]float64
	_, err = c.do(req, &resp)
	return resp, err
}
