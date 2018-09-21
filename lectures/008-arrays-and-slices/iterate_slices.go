package main

import (
  "fmt"
)

func doubleIt(x *int) {
  *x = *x + *x
}

func main() {
  s1 := []int { 2, 3, 5, 9 }
  for i := 0 ; i < len(s1) ; i++ { // classic C style
    fmt.Println("s1:", i, s1[i])
    doubleIt(&s1[i]) // original slice content modified
  }
  fmt.Println("Doubled s1", s1)
  fmt.Println("")

  s2 := []int { 4, 9, 11, 27 }
  for i := range s2 { // range of indexes
    fmt.Println("s2:", i, s2[i])
    doubleIt(&s2[i]) // original slice content modified
  }
  fmt.Println("Doubled s2", s2)
  fmt.Println("")

  s3 := []int { 12, 13, 15, 19 }
  for idx, value := range s3 { // range of indexes and values
    fmt.Println("s3:", idx, value)
    doubleIt(&s3[idx]) // original slice content modified
  }
  fmt.Println("Doubled s3", s3)
  fmt.Println("")

  s4 := []int { 72, 63, 55, 49 }
  sum := 0.0
  for _, value := range s4 { // range of ignored indexes and values, can't modify slice content
    sum += float64(value)
  }
  fmt.Println("Average of s4", sum/4.0)
  fmt.Println("")
}















//
