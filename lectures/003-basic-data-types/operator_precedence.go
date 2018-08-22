package main

import "fmt"

func main() {
  var z int32 = 3 + 4 * 5
  fmt.Println("z evaluates to = ", z)

  // We will change the precedence by adding paranthesis ()
  z = (3 + 4) * 5
  fmt.Println("After adding paranthesis z evaluates to = ", z)
}
