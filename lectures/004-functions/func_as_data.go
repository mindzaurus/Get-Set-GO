package main

import "fmt"

// x is value to increment, i is by how much to increment
func increment (x int, i int) int {
  return x + i
}

func doubleIncrement (x int, i int) int {
  return x * i
}

func addOne( x int, /* function as data -> 2nd argument */ fn func(x int, i int) int  ) {
  v := fn(x, 1) // fn is increment function
  fmt.Println("v:", v)
}

func addTwo( x int, fn func(x int, i int) int  ) {
  v := fn(x, 2)
  fmt.Println("v:", v)
}

func main() {
  addOne(10, increment) // fn is increment function
  addTwo(35, increment) // fn is increment function
  addTwo(35, doubleIncrement) // fn is increment doubleIncrement
}
