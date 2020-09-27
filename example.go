package main

import (
	"fmt"
	"github.com/prosline/httpclient/gohttp"
	"io/ioutil"
	"net/http"
)

var (
	c = getExampleClient()
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getExampleClient() gohttp.HttpClient {
	c := gohttp.New()
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC 123")
	c.SetHeaders(headers)
	return c
}

// Main example function
func main() {
	getURL()
}

// Get request functionality
func getURL() {
	res, err := c.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
}
func createUser(user User) {
	res, err := c.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
}
