package main

import (
	"fmt"
	"log"
	"net/http"
)

func StartAdminInterface(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("starting admin interface at http://%s\n", addr)
	admin := http.NewServeMux()
	admin.HandleFunc("/assets/css/", CssAssetHandler)
  admin.HandleFunc("/assets/js/", JsAssetHandler)
  admin.HandleFunc("/assets/fonts/", FontsAssetHandler)
	admin.HandleFunc("/add/", AddHandler)
	admin.HandleFunc("/delete/", DeleteHandler)
	admin.HandleFunc("/", HomeHandler)
	err := http.ListenAndServe(addr, admin)
	if err != nil {
		log.Printf("Error: ", err)
	}
}
