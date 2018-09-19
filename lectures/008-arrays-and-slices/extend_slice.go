package main

import (
  "fmt"
)

func main() {
  A := [6]int { 0, 1, 2, 3, 4, 5 }

  slice := A[1 : 4]
  fmt.Println(slice)
  fmt.Println("len(slice)", len(slice))
  fmt.Println("cap(slice)", cap(slice))
  fmt.Println("")

  slice2 := slice[0 : 5] // 1, 2, 3, 4, 5
  fmt.Println(slice2)
  fmt.Println("len(slice2)", len(slice2))
  fmt.Println("cap(slice2)", cap(slice2))
  fmt.Println("")

  slice = slice[0 : 4] // slice is modified using its own slice start, end, len, cap parameters
  fmt.Println(slice)
  fmt.Println("len(slice)", len(slice))
  fmt.Println("cap(slice)", cap(slice))
  fmt.Println("")

  slice = slice[3 : 4] // slice is modified using its own slice start, end, len, cap parameters
  fmt.Println(slice)
  fmt.Println("len(slice)", len(slice))
  fmt.Println("cap(slice)", cap(slice))
  fmt.Println("")
  
  slice = A[1 : 4]
  fmt.Println(slice)
  fmt.Println("len(slice)", len(slice))
  fmt.Println("cap(slice)", cap(slice))


}
