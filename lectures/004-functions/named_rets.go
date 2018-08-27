package main

import "fmt"

// Function that returns squared and cubed values
func sqAndCube(x int) /* start named args */ (sq int, cu int) /* end named args */  {
  sq = x * x
  cu = x * x * x
  return // plain return used not return sq, cu
}

func main() {
  s, c := sqAndCube(5)
  fmt.Println("Squared value:", s, "Cubed value:", c)
}
