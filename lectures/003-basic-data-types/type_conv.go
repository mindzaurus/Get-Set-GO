package main

import "fmt"

func main() {
  var f32 float32 = 0
  var pi float32 = 3.14

  var numDays int32 = 365
  var i int32 = -1

  i = int32(pi) // int32() is type casting / type conversion
  fmt.Println("float32 to int32", i)

  f32 = float32(numDays) // float32 is type casting / type conversion
  fmt.Println("int32 to float32", f32)
}
