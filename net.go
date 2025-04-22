package main

import (
	"net/http"
	"context"
	"io"
	"strings"
)


func CheckHTTPPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	return url
}

type AuthenticatedClient struct {
	client *http.Client
	token  string
	baseUrl string
}

type authTransport struct {
	token     string
	transport http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.token)
	return t.transport.RoundTrip(req)
}

func NewAuthenticatedClient(token string, baseUrl string) *AuthenticatedClient {
	return &AuthenticatedClient{
		client: &http.Client{
			Transport: &authTransport{
				token:     token,
				transport: http.DefaultTransport,
			},
		},
		token: token,
		baseUrl: baseUrl,
	}
}

func (c *AuthenticatedClient) Get(ctx context.Context, url string, data io.Reader) (*http.Response, error) {
	fullUrl := CheckHTTPPrefix(c.baseUrl) + url
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullUrl, data)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}
