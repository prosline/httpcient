package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)
const (
	DEFAULT_MAX_IDLE_CONNECTIONS = 5
	DEFAULT_CONNECTION_TIMEOUT = 2 * time.Second
	DEFAULT_RESPONSE_TIMEOUT = 4 * time.Second
)

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}
func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*Response, error) {
	allHeaders := c.getRequestHeaders(headers)
	reqBody, err := c.getRequestBody(allHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, errors.New("Can not create new request")

	}
	request.Header = allHeaders
	client := c.getHttpClient()
	// calling Go http.Client library
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	finalResponse := &Response{
		status:     response.Status,
		statusCode: response.StatusCode,
		headers:    response.Header,
		body:       responseBody,
	}

	return finalResponse, nil
}
func (c *httpClient)getHttpClient() *http.Client {

	c.clientOnce.Do(func() {
		c.client = &http.Client{
			Timeout: c.getConnectionTimeOut() + c.getResponseTimeOut(),
			Transport: &http.Transport{
				MaxConnsPerHost:       c.getMaxIdleconnections(),
				ResponseHeaderTimeout: c.getResponseTimeOut(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeOut(),
				}).DialContext,
			},
		}
	})
	return c.client
}

func (c *httpClient) getMaxIdleconnections() int{
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}
	return DEFAULT_MAX_IDLE_CONNECTIONS
}

func (c *httpClient) getResponseTimeOut() time.Duration{
	if c.builder.responseTimeout > 0{
		return c.builder.responseTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return DEFAULT_RESPONSE_TIMEOUT
}

func (c *httpClient) getConnectionTimeOut() time.Duration{
	if c.builder.connectionTimeout > 0{
		return c.builder.connectionTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return DEFAULT_CONNECTION_TIMEOUT
}
func (c *httpClient) getRequestHeaders(headers http.Header) http.Header {
	result := make(http.Header)
	// Including common headers to Request
	for h, v := range c.builder.headers {
		if len(v) > 0 {
			result.Set(h, v[0])
		}
	}
	// Including custom headers to Request
	for h, v := range headers {
		if len(v) > 0 {
			result.Set(h, v[0])
		}

	}
	return result
}
