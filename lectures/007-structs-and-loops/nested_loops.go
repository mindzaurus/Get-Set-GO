package main

import (
  "fmt"
)

func loopsSet1() {
  for i :=0; i <= 2; i++ {
    for j := 0; j <= 1; j++ {
      fmt.Println("loopsSet1 i =", i, "j =",j)
    }
    fmt.Println("----------")
  }
}

func loopsSet2() {
  for h :=0; h <= 3; h++ {
    fmt.Println("**********")
    for i :=0; i <= 2; i++ {
      for j := 0; j <= 1; j++ {
        fmt.Println("loopsSet2 h =", h, "i =", i, "j =",j)
      }
    }
  }
}

func main() {
  loopsSet1()
  fmt.Println()
  loopsSet2()
}
