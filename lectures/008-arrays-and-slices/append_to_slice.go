package main

import (
  "fmt"
)

func main() {
  s1 := []int { 2, 5, 19, 33, 75 }

  fmt.Println(s1)
  fmt.Println("len(s1)", len(s1), "cap(s1)", cap(s1))
  s1 = append(s1, 1000, 2002, 3003, 4004) // append values to slice
  fmt.Println(s1)
  fmt.Println("len(s1)", len(s1), "cap(s1)", cap(s1))
  fmt.Println("")

  s2 := []int { -50000, -1000000, 399999 }
  s1 = append(s1, s2...) // append slice to another slice
  fmt.Println(s1)
  fmt.Println("len(s1)", len(s1), "cap(s1)", cap(s1))
  fmt.Println("")

  s3 := []int {} // empty slice
  fmt.Println(s3)
  fmt.Println("len(s3)", len(s3), "cap(s3)", cap(s3))
  s3 = append(s3, 335, 497)
  fmt.Println(s3)
  fmt.Println("len(s3)", len(s3), "cap(s3)", cap(s3))
 }







//
