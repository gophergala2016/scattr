package main

import(
  "bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type FinalInput struct {
	Url      string
	Response string
}



func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}


func scattrHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ParseForm failed: %s\n", err.Error())
	}
	node := GetScattrData()
	method := r.Method
	payload := r.Form.Encode() //curl or web request that hits 8080(default port)
	ch := make(chan string)
	var buffer bytes.Buffer
	buffer.WriteString("{\"Responses\":[ ")
	fmt.Fprintf(os.Stdout, "Starting to scatter requests...\n")
	for _, url := range node.OutUrls {
		go scattr(url+r.URL.Path, method, payload, ch)
		str := <-ch
		buffer.WriteString(str)
		buffer.WriteString(",")
	}
	input := buffer.String()
	input = input[:len(input)-1] + " ]}"
	ListOfResponses := createFilteringInput(input)
	OutputWriter(ListOfResponses, w)
}

func createFilteringInput(input string) []FinalInput {
	var list map[string]interface{}
	var finalResponse []interface{}
	var FinalInputVar []FinalInput
	if err := json.Unmarshal([]byte(input), &list); err != nil {
		panic(err)
	}
	for key, _ := range list {
		finalResponse = list[key].([]interface{})
		for _, messageVal := range finalResponse {
			switch item := messageVal.(type) {
			case map[string]interface{}:
				var url string
				var response string
				if url1, ok := item["URL"].(string); ok {
					url = url1
				} else {
					url = url1
				}
				if response1, ok2 := item["Response"].(string); ok2 {
					response = response1
				} else {
					response = response1
					byteString, _ := GetBytes(item["Response"])
					response = string(byteString)
				}
				input1 := FinalInput{Url: url, Response: response}
				FinalInputVar = append(FinalInputVar, input1)
			default:
				fmt.Println("unknown type")
			}
		}
	}
	return FinalInputVar
}

func OutputWriter(FinalInputVar []FinalInput, w http.ResponseWriter) {
	nodes := GetScattrData()
	output, _ := evaluateScript(nodes.Data, FinalInputVar)
	fmt.Fprintf(w, output)
}

func StartScattrInterface(host string, port int) {
  addr := fmt.Sprintf("%s:%d", host, port)
  log.Printf("starting scattr interface at http://%s\n", addr)
  scattr := http.NewServeMux()
	scattr.HandleFunc("/", scattrHandler)
  scattr.HandleFunc("/{path}", scattrHandler)
  err := http.ListenAndServe(addr, scattr)
  if err != nil {
    log.Printf("Error: ", err)
  }

}
