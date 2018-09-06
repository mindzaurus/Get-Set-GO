// We will look at what are memory locations

package main

import (
  "fmt"
)

func main() {

  var age uint8 = 26

  fmt.Println("age variable holds ", age)
  fmt.Println("Address of age variable ", &age)

  var agePointer *uint8 = nil // no valid address is stored in pointer now, its nil
  agePointer = &age // store the age variables address inside the agePointer variable

  fmt.Println("agePointer is variable whose address is ", &agePointer)
  fmt.Println("agePointer variable holds address ", agePointer)
  fmt.Println("Reading the value pointed by the address stored inside agePointer ", *agePointer)

  *agePointer += 3 // this will increment age variable by 3 since it writes to the address of age directly
  fmt.Println("age variable now holds ", age)
}
