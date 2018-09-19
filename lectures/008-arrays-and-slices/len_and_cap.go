package main

import (
  "fmt"
)

func main() {
  A := [6]int { 0, 1, 2, 3, 4, 5 }
  slice := A [ 1 : 4 ]
  
  fmt.Println("len(slice) =", len(slice))
  fmt.Println("cap(slice) =", cap(slice))

  fmt.Println("")
  fmt.Println(slice)
}
