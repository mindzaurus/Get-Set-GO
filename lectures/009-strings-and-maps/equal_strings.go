package main

import (
  "fmt"
)

func main() {

  a := "apple"
  _a := "apple"

  b := "Apple"
  c := "Peach"

  fmt.Println("a == _a", a == _a)
  fmt.Println("a == b", a == b)
  fmt.Println("a == c", a == c)
  fmt.Println("")
  fmt.Println("a != _a", a != _a)
  fmt.Println("a != b", a != b)
  fmt.Println("a != c", a != c)
}
