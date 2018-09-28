package main

import (
  "fmt"
)

func main() {

  days := map[string] int {
    "Jan": 31, "Feb": 28,
    "Mar": 31, "Apr": 30,
    "May": 31, "Jun": 30,
  }

  fmt.Println(days)

  for month, numDays := range days { // key => month, value => numDays
    fmt.Println("month => ", month, " numDays => ", numDays)
  }

  fmt.Println("")

  for month := range days { // Iterating only key => month
    fmt.Println("month => ", month, " days[month] => ", days[month])
  }
}





















//
