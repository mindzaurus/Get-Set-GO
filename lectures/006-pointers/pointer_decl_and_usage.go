
package main

import (
  "fmt"
)

func main() {

  var iVal1 int64 = 7363992
  fmt.Println("iVal1", iVal1)

  // We will declare a pointer
  var iPtr1 *int64

  // We will store the address of variable iVal1's memory location in iPtr1
  iPtr1 = &iVal1 // using & address operator on iVal1, = is used to store address inside a pointer
  fmt.Println("iPtr1", iPtr1)

  iPtr2 := &iVal1 // shorter form
  fmt.Println("iPtr2", iPtr2)

  var iPtr3 *int64 = new(int64) // new allocated memory to hold a int64 value
  fmt.Println("iPtr3", iPtr3)

  iPtr4 := new(float64) // new allocated memory to hold a int64 value
  fmt.Println("iPtr4", iPtr4)
}
