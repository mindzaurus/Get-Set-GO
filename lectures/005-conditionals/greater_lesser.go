package main

import (
  "fmt"
)

func greaterLesser(x int) {
  fmt.Println("Checking ", x)

  if x > 55 { // x greater than 55
    fmt.Println("x is greater than 55")
  } else {
    fmt.Println("x is lesser than or equal to 55")
  }

  if x < 2 { // x lesser than 2
    fmt.Println("x is lesser than 2")
  } else {
    fmt.Println("x is greater than 2 or equal to 2")
  }
}

func main() {
  greaterLesser(54)
  greaterLesser(55)
  greaterLesser(56)

  fmt.Println(); fmt.Println()

  greaterLesser(1)
  greaterLesser(2)
  greaterLesser(3)
}
