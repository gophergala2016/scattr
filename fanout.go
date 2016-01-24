package main

import (
	"fmt"
	"github.com/mreiferson/go-httpclient"
	"net/http"
	"strings"
	"time"
)

type JSONMsg struct {
	URL      string
	Response string
}

type ErrorMsg struct {
	URL      string
	Response string
}

type JSONResponse struct {
	URL      string
	Response string
}


func scattr(url string, method string, payload string, ch chan string) {
	transport := &httpclient.Transport{
		ResponseHeaderTimeout: 2 * time.Second,
	}
	defer transport.Close()
	req, _ := http.NewRequest(method, url, strings.NewReader(payload))
	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	respStr := createResponse(url, resp, err)
	fmt.Printf("response from %s is %s \n", url, respStr)
	ch <- respStr
}
