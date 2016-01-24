package main

import (
	"log"
	"strconv"
	"strings"
)

func deleteUrl(index string) {
  conf := GetScattrData()
	i, err := strconv.Atoi(index)
	if err != nil {
		log.Println("ERROR:", "strconv failed for ", index, err.Error())
	}
	conf.OutUrls = append(conf.OutUrls[:i], conf.OutUrls[i+1:]...)
	conf.Flush()
}

func addUrl(url string) {
	conf := GetScattrData()
	url = strings.TrimSpace(url)
	if len(url) > 0 {
		conf.OutUrls = append(conf.OutUrls, url)
	}
	conf.Flush()
}
