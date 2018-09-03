package main

import (
  "fmt"
  "math"
)

// x is numerator, y is denominator
func divideFloats(x, y float32) float32 {
  // Any non zero value divided by zero is infinity, we will ignore negative infinity for now
  if y == 0 { // paranthesis () not required for condition check y == 0 like c, c++ or java
    return float32(math.Inf(1)) // if condition y == 0 is true divideFloats returns from here
  } // end of if block

  return x/y // when y is not 0 this line return x/y is executed
}

func main() {
  fmt.Println("divideFloats(2.0, 0)", divideFloats(2.0, 0))
  fmt.Println("divideFloats(2.0, 6.0)", divideFloats(2.0, 6.0))
}
