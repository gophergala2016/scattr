package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func createResponse(url string, resp *http.Response, err error) string {
	var respString string
	if err != nil {
		respString = createErrorResponse(url, err)
	} else {
		var contentType string
		contentType = getContentType(resp)
		respString = getResponseType(contentType, url, resp)
	}
	return respString
}

func createErrorResponse(url string, err error) string {
	var msg string
	if strings.Contains(err.Error(), "timeout awaiting response headers") {
		msg = "RESPONSE HAS TIMED OUT"
	} else {
		msg = fmt.Sprintf("%v", err)
	}
	errMsg := &ErrorMsg{url, msg}
	m, _ := json.Marshal(errMsg)
	updatedMsg := string(m)
	return updatedMsg
}

func getContentType(resp *http.Response) string {
	var contentType string
	if strings.Contains(resp.Header["Content-Type"][0], "text/html") {
		contentType = "html"
	} else if strings.Contains(resp.Header["Content-Type"][0], "application/json") {
		contentType = "json"
	} else if strings.Contains(resp.Header["Content-Type"][0], "application/octet-stream") {
		contentType = "binary"
	}
	return contentType
}

func getResponseType(contentType string, url string, resp *http.Response) string {
	var respStr string
	switch contentType {
	case "html":
		respStr = createHTMLResponse(url, resp)
	case "json":
		respStr = createJSONResponse(url, resp)
	case "binary":
		respStr = createBinaryResponse(url, resp)
	}
	return respStr
}

func createJSONResponse(url string, resp *http.Response) string {
	msg := readResponse(url, resp)
	jsonMsg := JSONMsg{url, msg}
	m, _ := json.Marshal(jsonMsg)
	updatedMsg := string(m)
	return updatedMsg
}

func createHTMLResponse(url string, resp *http.Response) string {
	msg := readResponse(url, resp)
	j := JSONResponse{URL: url, Response: msg}
	b, _ := json.Marshal(j)
	return string(b[:])
}

func createBinaryResponse(url string, resp *http.Response) string {
	msg := readResponse(url, resp)
	str := base64.StdEncoding.EncodeToString([]byte(msg))
	return "{" + "\"URL\":" + "\"" + url + "\"" + "," + "\"Response\":" + "\"" + str + "\"" + "}"
}

func readResponse(url string, resp *http.Response) string {
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	msg := string(contents)
	return msg
}
