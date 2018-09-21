package main

import (
  "fmt"
)

func main() {
  s1 := make([]int, 6, 10) // len == 6, cap == 10
  s1[2] = 67
  fmt.Println("s1", s1)

  s2 := make([]int, 10) // len == cap == 10
  s2[3] = 20382
  fmt.Println("s2", s2)

  // make is hidden
  s3 := []int {} // empty slice no elements in it can add elements using append
  // s3 = append(s3, 39) // we will append in a detailed manner later
  fmt.Println("s3", s3)

  // make is hidden
  s4 := []int {4, 9, 15}
  fmt.Println("s4", s4)

  // We will allocate huge slice of 40GB haha!
  hugeSzMem := 40 * (1024 * 1024 * 1024) // total 40GB, (1024 * 1024 * 1024) is 1GB
  s5 := make([]int8, hugeSzMem) // note int8 not int
  s5[0] = 99 // making compiler happy and we will only print s5[0]
  fmt.Println("s5[0]", s5[0])
}





//
