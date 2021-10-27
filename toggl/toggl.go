/*
Package toggl is a library of Toggl API v8 for the Go programming language.

See API documentation for more details.
https://github.com/toggl/toggl_api_docs/blob/master/toggl_api.md
*/
package toggl

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL string = "https://api.track.toggl.com/"
	apiVersionPath string = "api/v8/"
)

// Client is a client for interacting with Toggl API v8.
type Client struct {
	url        *url.URL
	httpClient *http.Client

	apiToken string
}

// NewClient creates a new Toggl API v8 client.
func NewClient(options ...Option) *Client {
	newClient := &Client{
		httpClient: http.DefaultClient,
	}
	newClient.setBaseURL(defaultBaseURL)

	for _, option := range options {
		option.apply(newClient)
	}

	return newClient
}

func (c *Client) setBaseURL(baseURL string) {
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}
	url, _ := url.Parse(baseURL)
	if !strings.HasSuffix(url.Path, apiVersionPath) {
		url.Path += apiVersionPath
	}
	c.url = url
}

// Option is an option for a Toggl API v8 client.
type Option interface {
	apply(*Client)
}

// WithHTTPClient returns a Option that specifies the HTTP client for communication.
func WithHTTPClient(httpClient *http.Client) Option {
	return &httpClientOption{httpClient: httpClient}
}

type httpClientOption struct {
	httpClient *http.Client
}

func (h *httpClientOption) apply(c *Client) {
	c.httpClient = h.httpClient
}

// WithAPIToken returns a Option that specifies an API token for authentication.
func WithAPIToken(apiToken string) Option {
	return apiTokenOption(apiToken)
}

type apiTokenOption string

func (a apiTokenOption) apply(c *Client) {
	c.apiToken = string(a)
}

var (
	// ErrContextNotFound is returned when the provided context is nil.
	ErrContextNotFound = errors.New("the provided context must be non-nil")

	// ErrAuthenticationFailure is returned when the API returns 403.
	ErrAuthenticationFailure = errors.New("authentication failed")
)
