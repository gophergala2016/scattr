package main

import (
  "os"
  "log"
)

func SavingScript(content string) {
  f, err := os.OpenFile(*ScriptFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Println("ERROR:", err)
	}
	defer f.Close()
	f.WriteString(content)
	f.Sync()
}
