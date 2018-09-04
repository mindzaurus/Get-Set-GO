package main

import (
  "fmt"
)

func gtEqLeEq(x int) { // gtEqLeEq -> Greater than or Equal to lesser than or equal to
  fmt.Println("Checking ", x)

  if x >= 100 { // x greater than or equal to 100
    fmt.Println("x greater than or equal to 100")
  } else {
    fmt.Println("x is lesser than 100")
  }

  if x <= 30 { // x lesser than or equal to 30
    fmt.Println("x lesser than or equal to 30")
  } else {
    fmt.Println("x is greater than 30")
  }
}

func main() {
  gtEqLeEq(99)
  gtEqLeEq(100)
  gtEqLeEq(101)

  fmt.Println(); fmt.Println()

  gtEqLeEq(29)
  gtEqLeEq(30)
  gtEqLeEq(31)
}
