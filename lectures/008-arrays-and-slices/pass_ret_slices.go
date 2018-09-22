package main

import (
  "fmt"
)

var theSlice []int // just reference nil reference in the beginning

func slicePassByRefs(s []int) {
  s[0] *= -50
  s[1] *= 10000
}

func slicePassByRefDemo() {
  s1 := []int { 2, 5, 11, 29, 45 }

  fmt.Println("s1 before slicePassByRefs(s1) => ", s1)
  slicePassByRefs(s1)
  fmt.Println("s1 after slicePassByRefs(s1) => ", s1)
}

func sliceRetByRef() []int {
  return theSlice
}

func sliceRetByRefDemo() {
  A := [...]int { 12, 15, 111, 129, 145 } // A is an array
  fmt.Println("A before returnedSlice := sliceRetByRef() ", A)

  theSlice = A[0:3]
  returnedSlice := sliceRetByRef() // get the slice returned from sliceRetByRef
  fmt.Println("returnedSlice => ", returnedSlice)

  returnedSlice[0], returnedSlice[1] = 555555, 999999 // A will be modified
  fmt.Println("A after returnedSlice := sliceRetByRef() ", A)

}

func main() {
  slicePassByRefDemo()
  sliceRetByRefDemo()
}














//
