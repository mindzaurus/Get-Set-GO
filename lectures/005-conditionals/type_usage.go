package main

import (
  "fmt"
)

type (
  WholeNumber int
  AnotherWholeNumber int
)

type ui64 uint64

func main() {
  var wn WholeNumber = 23
  fmt.Println(wn)

  var awn AnotherWholeNumber = 45
  fmt.Println(awn)

  t1 := int(wn)
  t2 := int(awn)
  fmt.Println("t1", t1, "t2", t2)

  wn = WholeNumber(t1)
  fmt.Println("wn = WholeNumber(t1)", wn)

  // wn = awn
  wn = WholeNumber(awn)
  fmt.Println("wn = WholeNumber(awn)", wn)

  var wn2 WholeNumber = 74
  wn = wn2
  fmt.Println(wn)

  var u ui64 = 7382
  fmt.Println(u)

  // u = wn
  u = ui64(wn)
  fmt.Println("u = ui64(wn)", u)

}
