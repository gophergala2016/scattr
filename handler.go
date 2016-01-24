package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("base")
	template.Must(tmpl.Parse(BaseTmpl))
	template.Must(tmpl.Parse(ConfigTemp))
	nodes := GetScattrData()
	tmpl.Execute(w, nodes)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ParseForm failed: %s\n", err.Error())
	}
	url := r.FormValue("url")
	addUrl(url)
	http.Redirect(w, r, "/", http.StatusFound)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	deleteUrl(id)
	http.Redirect(w, r, "/", http.StatusFound)
}

func CssAssetHandler(w http.ResponseWriter, r *http.Request) {
	asset := r.URL.Path[1:]
	data, err := Asset(asset)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	w.Header().Set("Content-Type", "text/css")
	w.Write(data)
}

func JsAssetHandler(w http.ResponseWriter, r *http.Request) {
	asset := r.URL.Path[1:]
	data, err := Asset(asset)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(data)
}


func FontsAssetHandler(w http.ResponseWriter, r *http.Request) {
	asset := r.URL.Path[1:]
	data, err := Asset(asset)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	w.Header().Set("Content-Type", "font/opentype")
	w.Write(data)
}
