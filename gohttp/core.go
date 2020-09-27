package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
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
func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}
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
	return client.Do(request)

}
func (c *httpClient) getRequestHeaders(headers http.Header) http.Header {
	result := make(http.Header)
	// Including common headers to Request
	for h, v := range c.Headers {
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
