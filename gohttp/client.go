package gohttp

import (
	"net/http"
	"sync"
)

type httpClient struct {
	builder *clientBuilder

	client     *http.Client
	clientOnce sync.Once
}

type Client interface {

	// Client http calls
	Get(url string, header http.Header) (*Response, error)
	Post(url string, header http.Header, body interface{}) (*Response, error)
	Put(url string, header http.Header, body interface{}) (*Response, error)
	Patch(url string, header http.Header, body interface{}) (*Response, error)
	Delete(url string, header http.Header) (*Response, error)
}

/*
	Public interface to initialize the httpclient
*/

func (c *httpClient) Get(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPost, url, headers, body)

}
func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPut, url, headers, body)

}
func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPatch, url, headers, body)

}
func (c *httpClient) Delete(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)

}
