package main

import (
  "fmt"
  "log"
  "net/http"
)

func StartAdminInterface(host string, port int) {
  addr := fmt.Sprintf("%s:%d", host, port)
  log.Printf("starting admin interface at http://%s\n", addr)
  err := http.ListenAndServe(addr, nil)
  if err != nil {
    log.Printf("Error: ", err)
  }
}
