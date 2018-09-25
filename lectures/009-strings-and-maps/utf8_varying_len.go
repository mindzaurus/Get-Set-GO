package main

import (
  "fmt"
)

// Common functionality write functions
// DRY principle - don't repeat yourself
func printInfo(c string) {
  num := 1

  for i, r := range c { // r => rune
    fmt.Printf("%d -> starting index = %d, char value = %d, char = %q \n", num, i, r, r)
    num++
  }
}

func engCat() { // English
  printInfo("cat")
}

func korCat() { // Korean
  printInfo("고양이")
}

func main() {
  engCat()
  fmt.Println("")
  korCat()
}
