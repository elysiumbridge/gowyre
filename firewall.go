package gowyre

import (
	"context"
)

// FirewallPermit https://docs.sendwyre.com/docs/rate-limiting#firewall-permit
func (c *Client) FirewallPermit(ctx context.Context, ips []string) error {
	url := "/v3/firewall/permit"
	if len(ips) > 0 {
		url += "?ips=" + ips[0]
		for _, ip := range ips[1:] {
			url += "," + ip
		}
	}
	req, err := c.newRequest(ctx, "POST", url, nil)
	if err != nil {
		return err
	}
	var resp interface{}
	_, err = c.do(req, &resp)
	return err
}
