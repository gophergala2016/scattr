package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"bytes"
)

type ScattrData struct {
	OutUrls []string
	Data string
}

func NewConfigFromFile(fname string) (*ScattrData, error) {
	config := new(ScattrData)
	if _, err := toml.DecodeFile(fname, &config); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return config, nil
}

func GetScattrData() *ScattrData {
	conf, err := NewConfigFromFile(*ConfigFile)
	if err != nil {
		fmt.Println("ERROR:", err)
		return nil
	}
	fmt.Println("CONF ", conf)
	return conf
}

func (node *ScattrData) Flush() {
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
	log.Println("FILE: ", fileBuffer.String())
	f.WriteString(fileBuffer.String())
}
