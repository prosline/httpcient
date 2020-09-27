package gohttp

import (
	"net/http"
)

type httpClient struct {
	Headers http.Header
}

type HttpClient interface {
	SetHeaders(header http.Header)
	Get(url string, header http.Header) (*http.Response, error)
	Post(url string, header http.Header, body interface{}) (*http.Response, error)
	Put(url string, header http.Header, body interface{}) (*http.Response, error)
	Patch(url string, header http.Header, body interface{}) (*http.Response, error)
	Delete(url string, header http.Header) (*http.Response, error)
}

/*
	Public interface to initialize the httpclient
*/

func New() HttpClient {
	return &httpClient{}
}

func (c *httpClient) SetHeaders(header http.Header) {
	c.Headers = header
}
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
