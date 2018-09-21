package main

import (
  "fmt"
)

func retSlice(x int) []int {
  if x < 0 { // for negative values return nil slice
    var s1 []int
    return s1
  }
  return []int {1,2 } // successful
}

func main() {
  var s1 []int
  fmt.Println("s1", s1, "len(s1)", len(s1), "cap(s1)", cap(s1))

  s2 := []int {}
  fmt.Println("s2", s2, "len(s2)", len(s2), "cap(s2)", cap(s2))

  fmt.Println("s1 == nil", s1 == nil)
  fmt.Println("s2 == nil", s2 == nil)

  fmt.Println("retSlice(1) == nil", retSlice(1) == nil,  retSlice(1))
  fmt.Println("retSlice(-1) == nil", retSlice(-1) == nil, retSlice(-1))
}




//
