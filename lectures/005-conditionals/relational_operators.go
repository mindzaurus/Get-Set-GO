package main

import (
  "fmt"
)

func relOpAND(min, max, val int) {
  if (val >= min) && (val <=max) {
    fmt.Println("relOpAND:: ",val,"is within range")
  } else {
    fmt.Println("relOpAND:: ",val,"is out of range")
  }
}

func relOpOR(min, max, val int) {
  if (val < min) || (val > max) {
    fmt.Println("relOpOR:: ",val,"is out of range")
  } else {
    fmt.Println("relOpOR:: ",val,"is within range")
  }
}

func relNOT_OR(min, max, val int) {
  if !(val < min || val > max) {
    fmt.Println("relNOT_OR:: ",val,"is within range")
  } else {
    fmt.Println("relNOT_OR:: ",val,"is out of range")
  }
}

func main() {
  relOpAND(1, 100, 23)
  relOpAND(1, 100, 101)
  relOpAND(1, 100, 0)
  relOpAND(1, 100, 100)

  k := 10
  if ( (k > 2 && k% 5 == 0 ) || (k % 2 == 0) && (k %10 == 0)) {
    fmt.Println("k is either >2 and multiple of 5 or k is even and multiple of 10")
  }

  relOpOR(1, 100, 23)
  relOpOR(1, 100, 101)
  relOpOR(1, 100, 0)
  relOpOR(1, 100, 100)

  fmt.Println()
  relNOT_OR(1, 100, 23)
  relNOT_OR(1, 100, 101)
  relNOT_OR(1, 100, 0)
  relNOT_OR(1, 100, 100)

  /* NEWS for c, c++ programmers only boolean expressions in GO can be used */
  var b bool = true
  if b {
    fmt.Println("b is true")
  }

  b = false
  if !b {
    fmt.Println("b is false")
  }
}
