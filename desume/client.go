package desume

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

const (
	defaultBaseURL = "https://desu.win/manga/api/"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

type Params map[string]string

func (a Params) toQueryParams() url.Values {
	res := make(url.Values)
	for k, v := range a {
		res.Add(k, v)
	}
	return res
}

// Option представляет функцию для настройки клиента.
type Option func(*Client)

// WithBaseURL устанавливает базовый URL API.
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// NewClient создает новый экземпляр клиента API.
func NewClient(options ...Option) *Client {
	client := &Client{
		baseURL:    defaultBaseURL,
		httpClient: &http.Client{},
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func (c *Client) sendRequest(ctx context.Context, method string, endpoint string, params Params, body interface{}) (*http.Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, endpoint)

	if params != nil {
		u.RawQuery = params.toQueryParams().Encode()
	}

	var req *http.Request
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequestWithContext(ctx, method, u.String(), bytes.NewBuffer(jsonBody))
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequestWithContext(ctx, method, u.String(), nil)
		if err != nil {
			return nil, err
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
