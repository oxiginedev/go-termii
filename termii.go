package termii

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://api.termii.com.ng/api/"

type Option func(*Client)

type Client struct {
	c       *http.Client
	baseURL *url.URL

	Messaging *Messaging
}

func New(opts ...Option) (*Client, error) {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, err
	}

	if c.c == nil {
		c.c = http.DefaultClient
	}

	c.baseURL = baseURL
	c.Messaging = &Messaging{client: c}

	return c, nil
}

func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) {
		c.c = h
	}
}

func (c *Client) NewRequest(method, url string, payload interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	var buf *bytes.Buffer

	if payload != nil {
		if err := json.NewEncoder(buf).Encode(payload); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

type Response struct {
	*http.Response
}

func (c *Client) Do(ctx context.Context, req *http.Request, dest interface{}) (*Response, error) {
	if ctx == nil {
		return nil, errors.New("context must not be nil")
	}

	req = req.WithContext(ctx)

	res, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if code := res.StatusCode; code > http.StatusCreated {
		var s struct {
			Message string `json:"message"`
		}

		if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
			return nil, err
		}

		return nil, errors.New(s.Message)
	}

	if dest != nil {
		decErr := json.NewDecoder(res.Body).Decode(&dest)
		if decErr != nil && decErr != io.EOF {
			return nil, err
		}
	}

	return &Response{res}, nil
}
