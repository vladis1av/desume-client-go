package desume

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
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

// doRequest выполняет HTTP-запрос к API и возвращает ответ.
func (c *Client) doRequest(ctx context.Context, method string, endpoint string, params Params, body interface{}) (*http.Response, error) {
	req, err := c.newRequest(ctx, method, endpoint, params, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}

// newRequest создает новый HTTP-запрос с указанными параметрами.
func (c *Client) newRequest(ctx context.Context, method string, endpoint string, params Params, body interface{}) (*http.Request, error) {
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
	log.Print(u)
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
		// req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "*/*")

		// req.Header.Set("authority", "desu.win")
		// req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		// req.Header.Set("accept-language", "ru,en;q=0.9,ja;q=0.8")
		// req.Header.Set("cache-control", "max-age=0")
		// req.Header.Set("if-modified-since", "Thu, 21 Mar 2024 07:32:31 GMT")
		// req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "YaBrowser";v="24.1", "Yowser";v="2.5"`)
		// req.Header.Set("sec-ch-ua-mobile", "?0")
		// req.Header.Set("sec-ch-ua-platform", `"Windows"`)
		// req.Header.Set("sec-fetch-dest", "document")
		// req.Header.Set("sec-fetch-mode", "navigate")
		// req.Header.Set("sec-fetch-site", "none")
		// req.Header.Set("sec-fetch-user", "?1")
		// req.Header.Set("upgrade-insecure-requests", "1")
		req.Header.Set("User-Agent", "PostmanRuntime/7.37.0")
	} else {
		req, err = http.NewRequestWithContext(ctx, method, u.String(), nil)
		if err != nil {
			return nil, err
		}
	}

	return req, nil
}
