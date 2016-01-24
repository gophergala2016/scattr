package main

import (
	"os"
)

func (node *scattrData) Flush() {
	f, err := os.OpenFile(*ConfigFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, url := range node.OutUrls {
		f.WriteString(url + "\n")
	}
	f.Sync()
	// var fileBuffer bytes.Buffer
	// e := toml.NewEncoder(&fileBuffer)
	// err := e.Encode(node)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
