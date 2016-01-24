package main

import (
	"os"
	"bytes"
	"github.com/BurntSushi/toml"
	"log"
)

func (node *scattrData) Flush() {
	f, err := os.OpenFile(*ConfigFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var fileBuffer bytes.Buffer
	e := toml.NewEncoder(&fileBuffer)
	err1 := e.Encode(node)
	if err1 != nil {
		log.Fatal(err)
	}
	f.WriteString(fileBuffer.String())
}
