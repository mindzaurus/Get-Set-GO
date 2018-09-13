package main

import (
  "fmt"
)

type Elephant struct {
  origin string // African or Asian
}

func (e Elephant) eat (food string) {
  fmt.Println("I the", e.origin, "Elephant is eating ", food)
}

func main() {
  e1 := Elephant {"African"}
  e1.eat("grass")

  e2 := Elephant {"Asian"}
  e2.eat("Sugarcane")
}
