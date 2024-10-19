package httpclient

import (
	"net/http"
	"net/url"
)

type HttpClient struct {
	client         *http.Client
	defaultHeaders http.Header
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client:         &http.Client{},
		defaultHeaders: http.Header{},
	}
}

func (c *HttpClient) SetUserAgent(userAgent string) {
	c.defaultHeaders.Add("User-Agent", userAgent)
}

func (c *HttpClient) SetCookies(u *url.URL, cookies []*http.Cookie) {
	c.client.Jar.SetCookies(u, cookies)
}
