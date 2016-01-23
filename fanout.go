package main

import (
  "net/http"
  "fmt"
  "strings"
)

func scattr(url string, method string, payload string) {
    req, _ := http.NewRequest(method, url, strings.NewReader(payload))
    client := &http.Client{}
    resp, _ := client.Do(req)
    fmt.Println("RESP: ",resp)
  }
