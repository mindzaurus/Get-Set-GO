package main

import (
  "fmt"
)

func main() {

  a := "apple"
  b := "banana"

  _a := []byte(a)
  _b := []byte(b)

  // slice style concat of strings
  ts := append(_a, _b...)
  s := string(ts)
  fmt.Println("s := string(ts) => ", s)

  // use + operator on strings
  s = a + b
  fmt.Println("s = a + b => ", s)

  s = a + " " + b
  fmt.Println("s = a + \" \" + b => ", s)

  s = a + " and " + b + " and " + " peach " + " and " + " cherry "
  fmt.Println("long concat string", s)
}








//
