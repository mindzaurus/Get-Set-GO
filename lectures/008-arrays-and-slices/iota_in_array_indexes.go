package main

import (
  "fmt"
)

const (
  JAN = iota
  FEB
  MAR
  APR
  MAY
)

func main() {
  days := [...]int { JAN: 31, MAR: 31, MAY: 31 }
  fmt.Println(days)

  days[FEB] = 28
  days[APR] = 30

  fmt.Println(days)
}










//
