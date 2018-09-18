package main

import (
  "fmt"
)



func main() {

  A := [6]int { 98, 76, 23, 51, 60, 85 }
  fmt.Printf("type of array A => %T\n", A)
  fmt.Println(A)
  fmt.Println("")

  s1 := A[1 : 5] // 76, 23, 51, 60
  fmt.Printf("type of slice s1 => %T\n", s1)
  fmt.Println(s1)

  s2 := A[0 : 6] // 98, 76, 23, 51, 60, 85
  fmt.Printf("type of slice s2 => %T\n", s2)
  fmt.Println(s2)

  s3 := A[0 : 2] // 98, 76
  fmt.Printf("type of slice s3 => %T\n", s3)
  fmt.Println(s3)

  s4 := A[3 : 3] // => from 3 to 2 wrong!!
  fmt.Printf("type of slice s4 => %T\n", s4)
  fmt.Println(s4)

  s5 := A[3 : 4] // => from 3 to 3 just one element 51
  fmt.Printf("type of slice s5 => %T\n", s5)
  fmt.Println(s5)

  // s6 := A[3 : 6] // A has just 6 elements :-) we say 100th element ??
  // fmt.Printf("type of slice s6 => %T\n", s6)
  // fmt.Println(s6)
}











//
