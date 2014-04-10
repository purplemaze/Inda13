// Stefan Nilsson
// Daniel Cserhalmi 2014-04-10
//
// This program litsens on 1 or 3 different servers(depending on the method you use) that
// indicate the current (fake) temperature at KTH in Stockholm
// and prints the answer time and temp.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	server := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}
	for {
		before := time.Now()
		//res := Get(server[0])
		//res := Read(server[0], time.Second)
		res := MultiRead(server, time.Second)
		after := time.Now()
		fmt.Println("Response:", *res)
		fmt.Println("Time:", after.Sub(before))
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}
}

type Response struct {
	Body       string
	StatusCode int
}

// Get makes an HTTP Get request and returns an abbreviated response.
// Status code 200 means that the request was successful.
// The function returns &Response{"", 0} if the request fails
// and it blocks forever if the server doesn't respond.
func Get(url string) *Response {
	res, err := http.Get(url)
	if err != nil {
		return &Response{}
	}
	// res.Body != nil when err == nil
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}
	return &Response{string(body), res.StatusCode}
}

// README.md contains solutions to prev. bugs.
// Read makes an HTTP Get request to a specified url and returns 
// the response with status code 200 or 503 - service unavailable .
// If the server doesn't answer before the timeout, the response is 
// 504 - Gateway timeout
func Read(url string, timeout time.Duration) (res *Response) {
	response := make(chan *Response, 1) // creates a buffered channel
	go func() {
		response <-Get(url)
		
	}()
	select {
	case res = <-response: 
	case <-time.After(timeout):
		res = &Response{"Gateway timeout\n", 504}
	}
	return
}

// MultiRead makes an HTTP Get request to each url and returns
// the response of the first server to answer with status code 200.
// If none of the servers answer before timeout, the response is
// 503 – Service unavailable.
func MultiRead(urls []string, timeout time.Duration) (res *Response) {
	response := make(chan *Response, 3) // creates a buffered channel
	for _, url := range urls {
		go func(url string) {
		response <-Get(url)
		
		}(url)
	}
	select {
	case res = <-response: 
	case <-time.After(timeout):
		res = &Response{"Service unavailable - \n", 503}
	}
	return

}
