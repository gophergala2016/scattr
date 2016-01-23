package main

import (
  "net/http"
  "fmt"
  "strings"
  "encoding/json"
)

type JSONMsg struct {
	URL      string
	Response json.RawMessage
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
    req, _ := http.NewRequest(method, url, strings.NewReader(payload))
    client := &http.Client{}
    resp, err := client.Do(req)
    respStr := createResponse(url, resp, err)
  	fmt.Printf("response from %s is %s \n", url, respStr)
  	ch <- respStr
  }
