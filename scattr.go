package main

import(
  "fmt"
  "log"
  "net/http"
  "os"

)

func scattrHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(os.Stdout, "REQUEST:[%s]\n", r)
  r.ParseForm()
  node := GetScattrData() //reads the toml file for the list of fanout urls and returns the struct
  method := r.Method
  fmt.Println("REQUEST method: ", method)
  payload := r.Form.Encode()
  fmt.Println("REQUEST payload: ", payload)
  fmt.Println("FAN OUT URLS: ", node.outUrls)

}


func StartScattrInterface(host string, port int) {
  addr := fmt.Sprintf("%s:%d", host, port)
  log.Printf("starting scattr interface at http://%s\n", addr)
	http.HandleFunc("/", scattrHandler)
  err := http.ListenAndServe(addr, nil)
  if err != nil {
    log.Printf("Error: ", err)
  }

}
