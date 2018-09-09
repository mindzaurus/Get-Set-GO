package main

import (
  "fmt"
)

func add(a, b int) *int {
  other_b := b
  result := a + other_b
  fmt.Println("add: Address of local variable other_b", &other_b)
  fmt.Println("add: Address of local variable result", &result)
  return &result // perfectly safe to return local addresses in golang
}

func main() {
  v := add(90, 257)
  //fmt.Printf("%T \n", v)
  fmt.Println("v holds address", v)
  fmt.Println("Dereferencing v gives", *v)
}
