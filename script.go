package main

import (
)

func SavingScript(content string) {
  conf := GetScattrData()
  conf.Data = content
  conf.Flush()

}
