package main

import (
  "fmt"
)

func equalityChecks(x int) {
  fmt.Println("Testing x ", x)

  if x == 5 { // NOTE the == operator
    fmt.Println("Yes x is equal to 5")
  } else { // x is not 5
    fmt.Println("No x is not equal to 5")
  }

  if x != 6 { // NOTE the != operator
    fmt.Println("Yes x is not equal to 6")
  } else { // x is 6
    fmt.Println("No x is equal to 6")
  }

  fmt.Println();fmt.Println()
}

func main() {
  equalityChecks(5)
  equalityChecks(6)
}


//
