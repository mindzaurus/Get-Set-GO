// We will add these numbers and print the results
// 4 and 5, 3 and 9, 10 and 11, 33 and 91, 1024 and 100

package main

import "fmt"

func addInts(x int, y int) {
  var z int = x + y + 5
  fmt.Println("z is now set to adding x + y + 5 ", z)
}

func addIntsSilent(x int, y int) int {
  var z int = x + y + 5
  return z
}

func main() {
  addInts(4, 5)
  addInts(3, 9)
  addInts(10, 11)
  addInts(33, 91)
  addInts(1024, 100)

  v1 := addIntsSilent(200, 332)
  fmt.Println("v1 is now set to ", v1)
}
