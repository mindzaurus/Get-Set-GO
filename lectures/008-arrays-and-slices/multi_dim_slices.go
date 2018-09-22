package main

import (
  "fmt"
)

func twoDimSlice() {
  A := new([3][2]int)

  s1 := A[0:3]
  fmt.Println("s1 before setting values",s1)
  s1[0][0] = 100 // first element
  s1[0][1] = 200
  s1[1][0] = 300
  s1[1][1] = 400
  s1[2][0] = 500
  s1[2][1] = 600 // last element
  fmt.Println("s1 after setting values", s1)

  s2 := (new([3][2]int)) [0:3]
  fmt.Println("s2 before setting values", s2)
  s2[0][0] = 1100 // first element
  s2[0][1] = 1200
  s2[1][0] = 1300
  s2[1][1] = 1400
  s2[2][0] = 1500
  s2[2][1] = 1600 // last element
  fmt.Println("s2 before setting values", s2)
}

func threeDimSlice() {
  A := new([3][2][1]int)
  s1 := A[0:3]

  fmt.Println("s1 before setting values",s1)
  s1[0][0][0] = 100 // first element
  s1[0][1][0] = 200
  s1[1][0][0] = 300
  s1[1][1][0] = 400
  s1[2][0][0] = 500
  s1[2][1][0] = 600 // last element
  fmt.Println("s1 after setting values",s1)

  s2 := (new([3][2][1]int)) [0:3]
  fmt.Println("s2 before setting values", s2)
  s2[0][0][0] = 2100 // first element
  s2[0][1][0] = 2200
  s2[1][0][0] = 2300
  s2[1][1][0] = 2400
  s2[2][0][0] = 2500
  s2[2][1][0] = 2600 // last element
  fmt.Println("s2 after setting values", s2)
}

func main() {
  twoDimSlice()
  fmt.Println()
  threeDimSlice()
}
