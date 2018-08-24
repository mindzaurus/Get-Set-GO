package main

import "fmt"

// package level variable
var count int = 0
                      /* NO INPUT ARGUMENTS */
func incrementCount(                             ) /* NO return value*/ {
  count++
} // NO return value
                   /* NO INPUT ARGUMENTS */   /* return value type */
func currentCount(                          )        int        {
  return count
}

func main() {
  incrementCount() // No arguments, no return value
  c := currentCount() // No arguments but returns a value
  fmt.Println("currentCount is ", c) // expect 1

  incrementCount() // No arguments, no return value
  c = currentCount() // No arguments but returns a value
  fmt.Println("currentCount is ", c) // expect 2
}
