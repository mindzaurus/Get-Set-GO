package main

import "fmt"

func main() {
  var ZERO float32 = 0.0
  var ONE float32 = 1.0
  var NEGONE float32 = -1.0 // negative one

  var x float32 = ONE / ZERO
  var y float32 = ZERO / ZERO
  var z float32 = NEGONE / ZERO

  fmt.Println("ONE / ZERO = ", x)
  fmt.Println("NEGATIVE ONE / ZERO = ", z)
  fmt.Println("ZERO / ZERO = ", y)
}
