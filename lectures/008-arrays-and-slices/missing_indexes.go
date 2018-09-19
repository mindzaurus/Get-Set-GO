package main

import (
  "fmt"
)

func main() {

  A := [6]int { 0, 1, 2, 3, 4, 5 }

  s1 := A[ : 4 ] // == A[ 0 : 4 ] => 0, 1, 2, 3
  fmt.Println("s1 =>", s1)
  fmt.Println("A[ 0 : 4 ] =>", A[ 0 : 4 ])
  fmt.Println("")

  s2 := A[ 2 : ] // == A[ 2 : len(A) ] == A[ 2 : 6 ] => 2, 3, 4, 5
  fmt.Println("s2 =>", s2)
  fmt.Println("A[ 2 : len(A) ] =>", A[ 2 : len(A) ])
  fmt.Println("")

  s3 := A[ : ] // == A[ 0 : len(A) ] == A[ 0 : 6 ] => 0, 1, 2, 3, 4, 5
  fmt.Println("s3 =>", s3)
  fmt.Println("A[ 0 : len(A) ] =>", A[ 0 : len(A) ])
  fmt.Println("")
}
