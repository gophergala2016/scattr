package main

import (
  "os"
  "fmt"
  "bufio"
)

func GetScattrData() scattrData {
  d := scattrData{}
   d.Read()
  return d
}

func (data *scattrData) Read() {
  f,err := os.Open(*ConfigFile)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Unable to read config file: %s", err.Error())
  }
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    url := scanner.Text()
    data.outUrls = append(data.outUrls, url)
    if err := scanner.Err(); err != nil {
			break
		}
  }

}
