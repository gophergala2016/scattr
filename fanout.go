package main

import (
  "net/http"
  "fmt"
  "strings"
)

func scattr(url string, method string, payload string) {
    req, _ := http.NewRequest(method, url, strings.NewReader(payload))
    client := &http.Client{}
    resp, err := client.Do(req)
    respStr = createResponse(url, resp, err)
  	fmt.Printf("response from %s is %s \n", url, respStr)
  	ch <- respStr
  }
