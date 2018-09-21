package main

import (
  "fmt"
)

type Doubled struct {
  val      int
  twiceVal int
}

func main() {
  doubles := [2]Doubled { } // array

  doubles[0].val = 2
  doubles[0].twiceVal = 4

  doubles[1].val = 45
  doubles[1].twiceVal = 90

  fmt.Println("doubles => ", doubles)

  s1 := doubles[1:2] // slice
  fmt.Println("s1 => ",s1)

  s2 := []Doubled { {7, 14}, {33, 66}, {42, 84} }
  fmt.Println("s2 => ",s2)


  s3 := make([]Doubled, 4)

  s3[0].val = 101
  s3[0].twiceVal = 202

  s3[1].val = 1000
  s3[1].twiceVal = 2000

  s3[2].val = 81
  s3[2].twiceVal = 162

  s3[3].val = 67
  s3[3].twiceVal = 134

  fmt.Println("s3 => ",s3)
}















//
