package main

import (
  "fmt"
)

func variadicSum01(first int, rest ...int) int {
  fmt.Println("first:", first)
  fmt.Println("rest:", rest)

  sum := first
  for i := range rest {
    arg := rest[i]
    sum += arg
  }
  return sum
}

func variadicSum02(first int, second int, rest ...int) int {
  fmt.Println("first:", first)
  fmt.Println("second:", second)
  fmt.Println("rest:", rest)

  sum := first + second
  for i := range rest {
    arg := rest[i]
    sum += arg
  }
  return sum
}

func main() {
  r := 0
  r = variadicSum01(5, 23, 65)
  fmt.Println("variadicSum01(5, 23, 65)", r) ; fmt.Println("")
  r = variadicSum01(5, 23, 65, 129, 313, 456, 799, 812)
  fmt.Println("variadicSum01(5, 23, 65, 129, 313, 456, 799, 812)", r) ; fmt.Println("")

  r = variadicSum02(5, 23, 65)
  fmt.Println("variadicSum02(5, 23, 65)", r) ; fmt.Println("")
  r = variadicSum02(5, 23, 65, 129, 313, 456, 799, 812)
  fmt.Println("variadicSum02(5, 23, 65, 129, 313, 456, 799, 812)", r) ; fmt.Println("")

  slice := []int { 65, 129, 313, 456, 799, 812 } // this is slice
  r = variadicSum02(-5, -23, slice...)
  fmt.Println("variadicSum02(-5, -23, slice...)", r) ; fmt.Println("")
}
