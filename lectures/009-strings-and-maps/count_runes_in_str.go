package main

import (
  "fmt"
  "unicode/utf8"
)

func quoted(s string) string {
  return `"` + s + `"`
}

func printLengthInfo(s string) {
  fmt.Println("len(", quoted(s), ") =", len(s))
  fmt.Println("utf8.RuneCountInString(", quoted(s), ") =", utf8.RuneCountInString(s))
  fmt.Println("")
}

func main() {
  nonEng := "collège in Zürich" // non English
  eng := "college in Zurich" // English

  printLengthInfo(nonEng)
  printLengthInfo(eng)
}
