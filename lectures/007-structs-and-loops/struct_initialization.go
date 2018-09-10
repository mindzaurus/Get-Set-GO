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
  goBook := Product { 29786, "The GO book", BOOKS, 25.45 }

  fmt.Println(goBook)
  fmt.Println("goBook.ID", goBook.ID)
  fmt.Println("goBook.name", goBook.name)
  fmt.Println("goBook.category", goBook.category)
  fmt.Println("goBook.price", goBook.price)
}






//
