package main

import (
	"fmt"
	"github.com/prosline/httpclient/gohttp"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	b = getExampleClient()
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getExampleClient() gohttp.ClientBuilder {
	c := gohttp.NewBuilder()
	c.SetMaxIdleConnections(20)
	c.SetConnectionTimeout(2 * time.Second)
	c.SetResponseTimeout(4 * time.Second)
	c.DisableTimeOuts(true)
	headers := make(http.Header)
	// The statement below forces a bad request since "Bearer ABC 123" is an invalid token
	//headers.Set("Authorization", "Bearer ABC 123")
	c.SetHeaders(headers)
	return c
}

// Main example function
func main() {
	getURL()
}

// Get request functionality
func getURL() {
	c := b.Build()
	res, err := c.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
}
func createUser(user User) {
	c := b.Build()
	res, err := c.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
}
