package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

func NewConfigFromFile(fname string) (*scattrData, error) {
	config := new(scattrData)
	if _, err := toml.DecodeFile(fname, &config); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return config, nil
}

func GetScattrData() *scattrData {
	conf, err := NewConfigFromFile(*ConfigFile)
	if err != nil {
		fmt.Println("ERROR:", err)
		return nil
	}
	fmt.Println("CONF ", conf)
	return conf
}
