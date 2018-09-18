package main

import (
  "fmt"
)

func handleArray(ap *[3]int) {
  ap[0] += 100
  ap[1] += 100
  ap[2] += 100
}

func main() {
  var a2 [3]int = [3]int {3, 4, 5} // same as a2 := [3]int {3, 4, 5}
  fmt.Printf("a2 [3]int => type %T\n", a2)
  fmt.Println(a2)
  fmt.Println("")

  var a1 *[3]int = &[3]int {3, 4, 5} // same as a1 :=  &[3]int {3, 4, 5}
  fmt.Printf("a *[3]int => type %T\n", a1)
  fmt.Println("Before handleArray(a1)", a1)
  handleArray(a1)
  fmt.Println("After handleArray(a1)", a1)
  fmt.Println("")

  var b *[3]int = new ([3]int) // same as b := new ([3]int)
  b[0], b[1], b[2] = 303, 305, 309
  fmt.Println("Before handleArray(b)", b)
  handleArray(b)
  fmt.Println("After handleArray(b)", b)
}







//
