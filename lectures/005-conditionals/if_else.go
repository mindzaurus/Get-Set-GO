package main

import (
  "fmt"
)

// { and else should be in the same line
func isXGreater(x, y int) {
  if x > y { // testing is x greater than y
    fmt.Println("x=",x,"    y=",y, "x is greater than y")
  } else { // either x is equal to y or x is lesser than y
    fmt.Println("x=",x,"    y=",y, "x is equal to y or x is lesser than y")
  }
}

func main() {
  isXGreater(5, 1)
  isXGreater(5, 12)
  isXGreater(10, 10)
}
