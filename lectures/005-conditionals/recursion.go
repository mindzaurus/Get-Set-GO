package main

import (
  "fmt"
)

// printTill(5) will print from 1 to 5
func printTill(x int) {

  if x == 0 { // termination condition
    return
  }

  printTill( x - 1 ) // This is the actual recursive call with x - 1 argument

  fmt.Println(x) // Action to be performed after returning from recursion
}


func main() {
  printTill(5)
}
