package main

import (
  "fmt"
)

func loop1() {
  i :=  5

  for i <= 250 { // same as => for i := 5; i <= 250; i += 35 
    // start of code block
    fmt.Println("loop1 i =",i)
    i += 35 // this changes the condition per iteration
  } // start of code block
}

func loop2() {
  for i := 5; i <= 250; i += 35 {
    // start of code block
    fmt.Println("loop2 i =",i)
  } // start of code block
}

func main() {
  loop1()
  fmt.Println()
  loop2()
}
