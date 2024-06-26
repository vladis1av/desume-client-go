package desume

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

// defaultBaseURL constant defines the default API base URL that will be used by the client unless another URL is specified.
const (
	defaultBaseURL = "https://desu.win/manga/api/"
)

// Client structure represents the API client that will be used to send HTTP requests to the API.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// Params type represents the request parameters that will be passed to the API URL.
type Params map[string]string

// Option type represents a function that can be used to configure a Client instance
type Option func(*Client)

func (a Params) toQueryParams() url.Values {
	res := make(url.Values)
	for k, v := range a {
		res.Add(k, v)
	}
	return res
}

// WithBaseURL feature option sets the base API URL for the Client instance
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// NewClient creates a new instance of the API client.
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

// sendRequest sends an HTTP request to the API with the specified parameters
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
