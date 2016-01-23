package main

import (
"net/http"
"strings"
)

func createResponse(url string, resp *http.Response, err error) {
  var respString string
  if err != nil {
    respString = createErrorResponse(url, err)
  } else {
    var contentType, responseType string
    contentType = getContentType(resp)
    respString = getResponseType(contentType, url, resp)
  }
 return respString
}

func createErrorResponse(url string, err error) string{
  var msg string
  if strings.Contains(err.Error(),"timeout awaiting response headers") {
		msg = "RESPONSE HAS TIMED OUT"
	} else {
		msg = fmt.Sprintf("%v", err)
	}
	return msg
}

}

func getContentType(resp *http.Response) string {
  var contentType string
	if strings.Contains(resp.Header["Content-Type"][0], "text/html") {
		contentType = "html"
	}else if strings.Contains(resp.Header["Content-Type"][0], "application/json") {
    contentType = "json"
  }else if strings.Contains(resp.Header["Content-Type"][0], "application/octet-stream") {
    contentType = "binary"
  }
  return contentType
}


func getResponseType(contentType string, url string, resp *http.Response) {
  var respStr string
	switch contentType {
	case "html":
		respStr = createHTMLResponse(url, resp)
	case "json":
		respStr = createJSONResponse(url, resp)
	case "binary":
		respStr = createBinaryResponse(url, resp)
	}
	//fmt.Printf("response from %s is %s \n", url, respStr)
	//fmt.Println(respStr)
	return respStr

}
