package posthook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	defaultEndpoint  = "https://api.posthook.io/v1/hooks"
	apiKeyHeaderName = "X-API-Key"
)

// Client interfaces with the Posthook API: api.posthook.io
type Client struct {
	apiKey     string
	endpoint   string
	httpClient *http.Client
}

// Config lets users configure the postHook struct
type Config func(c *Client)

// WithEndpoint configures the API endpoint
func WithEndpoint(endpoint string) Config {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

// WithHTTPClient replaces the underlying http.DefaultClient
func WithHTTPClient(httpClient *http.Client) Config {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// New creates a new postHook client
func New(apiKey string, opts ...Config) *Client {
	client := &Client{
		apiKey:     apiKey,
		endpoint:   defaultEndpoint,
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(client)
	}
	return client
}

// Endpoint returns the currently configured endpoint
func (c *Client) Endpoint() string {
	return c.endpoint
}

// Schedule creates a new hook to schedule a time for Posthook to make a request to your application.
func (c *Client) Schedule(path string, at time.Time, data interface{}) (*Hook, error) {

	hook := Hook{
		Path:   path,
		PostAt: at,
		Data:   data,
	}
	buf, err := json.Marshal(hook)
	if err != nil {
		return nil, err
	}
	httpReq, _ := http.NewRequest(http.MethodPost, c.endpoint, bytes.NewBuffer(buf))

	body, err := c.send(httpReq)
	if err != nil {
		return nil, err
	}
	return single(body)
}

// List retrieves Hooks that have already been scheduled.
func (c *Client) List(filters Filters) ([]Hook, error) {
	v, err := query.Values(filters)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(http.MethodGet, c.endpoint, nil)
	if err != nil {
		return nil, err
	}

	httpReq.URL.RawQuery = v.Encode()
	body, err := c.send(httpReq)
	if err != nil {
		return nil, err
	}

	return list(body)
}

// Get retrieves a specific Hook.
func (c *Client) Get(id string) (*Hook, error) {
	url := fmt.Sprintf("%s/%s", c.endpoint, id)
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.send(httpReq)
	if err != nil {
		return nil, err
	}

	return single(body)
}

// Delete a Hook. This will stop Posthook from making the request back to your application for this Hook.
func (c *Client) Delete(id string) error {
	url := fmt.Sprintf("%s/%s", c.endpoint, id)
	httpReq, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	_, err = c.send(httpReq)
	return err
}

func (c *Client) send(req *http.Request) ([]byte, error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(apiKeyHeaderName, c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("api returned http %d with body: %s", resp.StatusCode, content)
	}

	return content, err
}
