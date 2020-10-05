package main

import (
	"fmt"
	"github.com/prosline/httpclient/gohttp"
	"net/http"
	"time"
)

var (
	gitClient = getExampleClient()
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getExampleClient() gohttp.Client {
	headers := make(http.Header)
	c := gohttp.NewBuilder().SetMaxIdleConnections(20).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(4 * time.Second).
		DisableTimeOuts(true).
		SetHeaders(headers).Build()
	return c
}

// Main example function
func main() {
	for i := 0; i < 10; i++ {
		go func() {
			getURL()
		}()
	}
	time.Sleep(5 * time.Second)
}

// Get request functionality
func getURL() {
	res, err := gitClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Status())
	fmt.Println(res.StatusCode())
	fmt.Println(res.String())
	//	var user User
	//	if err := res.Unmarshal(&user); err != nil{
	//		panic(err)
	//	}
}
