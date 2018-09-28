package main

import (
  "fmt"
)

func bumpPrice( m map[string] float32, foodName string ) {
  m[foodName] += 1.0
}

func main() {
  fpm := make (map[string] float32) // key => string, value => float32

  fpm["Burrito"] = 4.65 // 4.65 dollars insert your currency here
  fpm["Sushi"] = 4.50

  fmt.Println("fpm[\"Sushi\"]", fpm["Sushi"])
  fmt.Println(fpm)

  bumpPrice(fpm, "Burrito") // pass by reference
  fmt.Println("fpm[\"Burrito\"]", fpm["Burrito"])
}










//
