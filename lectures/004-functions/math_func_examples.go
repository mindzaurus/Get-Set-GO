
package main

import (
  "math"
  "fmt"
)

func main() {
  // math.Pow
  fmt.Println("math.Pow(2.0, 5.0)", math.Pow(2.0, 5.0))

  // math.Sqrt
  fmt.Println("math.Sqrt(225.0)", math.Sqrt(225.0))

  // math.Exp
  fmt.Println("math.Exp(4.0)", math.Exp(4.0))

  // math.Log
  fmt.Println("math.Log(4.0)", math.Log(4.0))

  // Exp and Log are inverse functions like + and - are inverse functions
  fmt.Println("math.Log(math.Exp(4.0))", math.Log(math.Exp(4.0))) //should print 4.0
}
