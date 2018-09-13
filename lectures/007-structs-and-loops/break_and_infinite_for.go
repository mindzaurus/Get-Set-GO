package main

import (
  "fmt"
)

func loop1() {
  i := 0

  for { // this will loop for ever like infinite loop
    fmt.Println("loop1 i =",i)
    i++

    if i == 10 {
      break // this breaks the infinite loop
    }
  }
  fmt.Println("loop1: After for loop")
}

func loop2() {
  i := 0

  for { // this will loop for ever like infinite loop
    fmt.Println("loop2 i =",i)
    i++

    if i == 15 {
      return // jumps out of the loop2 function back to the caller
    }
  }
  fmt.Println("loop2: After for loop")
}

func main() {
  loop1()
  fmt.Println("")
  loop2()
}
