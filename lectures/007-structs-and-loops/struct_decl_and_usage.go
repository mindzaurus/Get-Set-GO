package main

import (
  "fmt"
)

type ProductCategory int
const (
  BOOKS ProductCategory = iota
  TOYS
  GROCERIES
)

type Product struct {
  ID       int
  name     string
  category ProductCategory
  price    float32
}

func main() {
  var marioToy Product
  marioToy.ID = 73820
  marioToy.name = "Mario The Plumber"
  marioToy.category = TOYS
  marioToy.price = 9.99

  fmt.Println("marioToy.ID", marioToy.ID)
  fmt.Println("marioToy.name", marioToy.name)
  fmt.Println("marioToy.category", marioToy.category)
  fmt.Println("marioToy.price", marioToy.price)
}

















//
