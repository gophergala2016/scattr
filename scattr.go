package main

import(
  "fmt"
  "log"
  "net/http"

)

func StartScattrInterface(host string, port int) {
  addr := fmt.Sprintf("%s:%d", host, port)
  log.Printf("starting scattr interface at http://%s\n", addr)
  err := http.ListenAndServe(addr, nil)
  if err != nil {
    log.Printf("Error: ", err)
  }

}
