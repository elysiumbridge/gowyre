package gowyre

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	testWyreURL = "https://api.testwyre.com"
	sendWyreURL = "https://api.sendwyre.com"
)

type Client struct {
	Secret    string `json:"secret"`
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

type environment string

var (
	Sandbox    environment = "test"
	Production environment = "production"
)

func NewClient(secret, agent string, env environment, httpClient *http.Client) (c *Client, err error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	var URL *url.URL
	if env == Sandbox {
		URL, err = url.Parse(testWyreURL)
	} else {
		URL, err = url.Parse(sendWyreURL)
	}
	if err != nil {
		return
	}

	userAgent := "gowyre"
	if agent != "" {
		userAgent = agent
	}

	c = &Client{secret, URL, userAgent, httpClient}
	return
}

func (c *Client) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL.String(), path), buf)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", "Bearer "+c.Secret)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		var werr WyreError
		if err = json.NewDecoder(resp.Body).Decode(&werr); err != nil {
			return nil, err
		}
		return nil, NewError(werr)
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
