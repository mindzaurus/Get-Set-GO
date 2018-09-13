package main

import (
  "fmt"
)

func loop1() {
  for i := 0; i <= 10; i++ {
    fmt.Println("loop1: i =", i)
  }
}

func loop2() {
  for i := 0; i <= 10; i++ {
    if i == 5 || i == 6 {
      continue // iteration continues with i == 6 below print is not printed for 5
    }
    fmt.Println("loop2: i =", i)
  }
}

func main() {
  loop1()
  fmt.Println()
  loop2()
}
