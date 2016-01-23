package main

import(
  "fmt"
  "log"
  "net/http"
  "os"
  "bytes"
)

func scattrHandler(w http.ResponseWriter, r *http.Request) {
  var buffer bytes.Buffer
  ch := make(chan string)
  var respString string
  fmt.Fprintf(os.Stdout, "REQUEST:[%s]\n", r)
  r.ParseForm()
  node := GetScattrData() //reads the toml file for the list of fanout urls and returns the struct
  method := r.Method
  payload := r.Form.Encode()
  buffer.WriteString("\"Responses\":[ ")
  for _, url := range node.outUrls {
		fmt.Fprintf(os.Stdout, "Calling fanout with %s, %s, %s\n", url+r.URL.Path, method, payload)
		go scattr(url+r.URL.Path, method, payload, ch)
    respString = <- ch
    log.Printf("Response: ", respString)

}
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
