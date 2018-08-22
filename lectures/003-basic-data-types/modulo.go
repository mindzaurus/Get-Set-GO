package main

import "fmt"

func main() {
  var i int32 = -1

  i = 42 % 8
  fmt.Println(" 42 modulo 8 is = ", i)

  i = 77 % 10
  fmt.Println(" 77 modulo 10 is = ", i)

  i = 77 % 11
  fmt.Println(" 77 modulo 11 is = ", i)

  i = 24 % 8
  fmt.Println(" 24 modulo 8 is = ", i)
}
