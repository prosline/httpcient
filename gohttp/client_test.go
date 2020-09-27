package gohttp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	//  Initialize
	client := httpClient{}
	commonHeader := make(http.Header)
	commonHeader.Set("Content-Type", "application/json")
	commonHeader.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeader

	//  Execute
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	//  Validate
	if len(finalHeaders) != 3 {
		t.Error("We expect 3 headers")
	}
	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("We expect value ABD-123")
	}
	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("We expect value ABD-123")
	}
	if finalHeaders.Get("User-Agent") != "http-client" {
		t.Error("We expect value cool-http-client")
	}
}
func TestGetRequest(t *testing.T) {
	// Initialization
	httpClient := httpClient{}
	t.Run("NoBodyNilResponse", func(t *testing.T) {
		reqBody, err := httpClient.getRequestBody("", nil)
		if err != nil {
			t.Error("Invalid body length! We expect body nil")
		}

		if reqBody != nil {
			t.Error("Invalid body length! We expect context type nil")
		}
	})

	t.Run("BodyJsonResponse", func(t *testing.T) {
		body := []string{"response", "json", "test"}
		reqBody, err := httpClient.getRequestBody("application/json", body)
		if err != nil {
			t.Error("Problem occurred while marshaling slice as json")
		}
		if string(reqBody) != `["response","json","test"]` {
			t.Error("Invalid Json data expecting body = [response,json,test]")
		}
	})
	t.Run("BodyXMLResponse", func(t *testing.T) {
		body := []string{"response", "json", "test"}
		reqBody, err := httpClient.getRequestBody("application/xml", body)
		fmt.Print(string(reqBody))
		if err != nil {
			t.Error("Problem occurred while marshaling slice as json")
		}
		if string(reqBody) != `<string>response</string><string>json</string><string>test</string>` {
			t.Error("Invalid XML data expecting body = <string>response</string><string>json</string><string>test</string>`")
		}
	})
}
