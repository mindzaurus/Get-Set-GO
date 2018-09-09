package main

import (
  "fmt"
)

func main() {
  var ptr *int = nil // initialize pointer to a known bad address (uninitialized address)

  if ptr == nil { // check whether pointer is nil, this operation is safe won't cause abort
    fmt.Println("pointer ", ptr, "is nil")
  }

  fmt.Println("Dereferencing nil pointer ", *ptr)
}
