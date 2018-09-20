package main

import (
  "fmt"
)

func main() {
  A := [6]int { 0, 1, 2, 3, 4, 5 }
  fmt.Println("A =>", A)

  slice1 := A[1:4] // => 1, 2, 3,
  fmt.Println("slice1 =>",slice1)

  slice2 := slice1[1:5]
  fmt.Println("slice2 before *100", slice2) // => 2, 3, 4, 5

  slice2[0] *= 100
  slice2[1] *= 100
  slice2[2] *= 100
  slice2[3] *= 100
  fmt.Println("slice2 after *100", slice2)

  // Changing elements of slice2 will change array A
  fmt.Println("after *100 on slice2 A => ", A)
}




//
