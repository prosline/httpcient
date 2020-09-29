package gohttp

import (
	"net/http"
	"time"
)
type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
	baseUrl            string
	client             *http.Client
	userAgent          string
}

type ClientBuilder interface{
	// Client Configuration
	DisableTimeOuts(disable bool) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(conn int) ClientBuilder
	SetHeaders(header http.Header) ClientBuilder
	SetUserAgent(userAgent string) ClientBuilder
	Build() HttpClient
}

func NewBuilder() ClientBuilder{
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() HttpClient{
	client := httpClient{
		builder: c,
	}
	return &client
}

func (c *clientBuilder) SetHeaders(header http.Header) ClientBuilder{
	c.headers = header
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder{
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder{
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(conn int) ClientBuilder{
	c.maxIdleConnections = conn
	return c
}

func (c *clientBuilder) DisableTimeOuts(disable bool) ClientBuilder{
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	c.userAgent = userAgent
	return c
}