package main

import (
  "fmt"
)

var childAge int = 8
var middleAge int = 40
var oldAge int = 95

func main() {

  var agePtr *int

  agePtr = &childAge
  fmt.Println("agePtr points to address &childAge", agePtr)
  fmt.Println("after agePtr = &childAge, *agePtr=", *agePtr) // *agePtr is reading contents of memory address pointed by agePtr

  agePtr = &middleAge
  fmt.Println("agePtr points to address &middleAge", agePtr)
  fmt.Println("after agePtr = &middleAge, *agePtr=", *agePtr) // now agePtr points to the address of middleAge

  agePtr = &oldAge
  fmt.Println("agePtr points to address &oldAge", agePtr)
  fmt.Println("after agePtr = &oldAge, *agePtr=", *agePtr) // // now agePtr points to the address of oldAge
}
