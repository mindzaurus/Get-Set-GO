package main

import (
  "fmt"
)

func loop1() {
  i, x := 0, 0

  for i = 0 ; i <= 5 ; i++ {
    fmt.Println("loop1 i =",i, "and x =",x)
    x += 2
  }
}

func loop2() {
  for i, x := 0, 0 ; i <= 10 ; i++ {
    fmt.Println("loop2 i =",i, "and x =",x)
    x += 3
  }
}

func main() {
  loop1()
  fmt.Println("")
  loop2()
}

//
