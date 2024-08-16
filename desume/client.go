package desume

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"time"

	"golang.org/x/time/rate"
)

// defaultBaseURL constant defines the default API base URL that will be used by the client unless another URL is specified.
const (
	defaultBaseURL = "https://desu.win/manga/api/"
)

var (
	ErrRateLimitExceeded = errors.New("rate limit exceeded")
)

// Client structure represents the API client that will be used to send HTTP requests to the API.
type Client struct {
	baseURL     string
	httpClient  *http.Client
	rateLimiter *rate.Limiter
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

// WithTimeout feature option sets the timeout for the HTTP client
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithMaxIdleConns feature option sets the maximum number of idle connections for the HTTP client
func WithMaxIdleConns(maxIdleConns int) Option {
	return func(c *Client) {
		c.httpClient.Transport.(*http.Transport).MaxIdleConns = maxIdleConns
	}
}

// WithIdleConnTimeout feature option sets the idle connection timeout for the HTTP client
func WithIdleConnTimeout(idleConnTimeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Transport.(*http.Transport).IdleConnTimeout = idleConnTimeout
	}
}

// WithTLSHandshakeTimeout feature option sets the TLS handshake timeout for the HTTP client
func WithTLSHandshakeTimeout(tlsHandshakeTimeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Transport.(*http.Transport).TLSHandshakeTimeout = tlsHandshakeTimeout
	}
}

// WithDisableCompression feature option sets the disable compression flag for the HTTP client
func WithDisableCompression(disableCompression bool) Option {
	return func(c *Client) {
		c.httpClient.Transport.(*http.Transport).DisableCompression = disableCompression
	}
}

// WithRateLimiter feature option sets the rate limiter for the Client instance
func WithRateLimiter(rps int, burst int) Option {
	return func(c *Client) {
		c.rateLimiter = rate.NewLimiter(rate.Limit(rps), burst)
	}
}

// NewClient creates a new DESUME Client instance
func NewClient(options ...Option) *Client {
	httpClient := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        20,
			IdleConnTimeout:     60 * time.Second,
			TLSHandshakeTimeout: 5 * time.Second,
			DisableCompression:  true,
		},
	}

	client := &Client{
		baseURL:    defaultBaseURL,
		httpClient: httpClient,
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

	if c.rateLimiter != nil {
		if !c.rateLimiter.Allow() {
			return nil, ErrRateLimitExceeded
		}
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
		req.Header.Set("Content-Type", "application/json")
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
