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

type HttpClient interface {

	// Client http calls
	Get(url string, header http.Header) (*http.Response, error)
	Post(url string, header http.Header, body interface{}) (*http.Response, error)
	Put(url string, header http.Header, body interface{}) (*http.Response, error)
	Patch(url string, header http.Header, body interface{}) (*http.Response, error)
	Delete(url string, header http.Header) (*http.Response, error)
}

/*
	Public interface to initialize the httpclient
*/

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)

}
func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)

}
func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)

}
func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)

}
