package main

import (
  "fmt"
)

// return one more than given value v
func oneMore(v int) int {
  return v + 1
}

func ifCheck(x int) {
  if t1 := oneMore(x); t1 == 1 {
    fmt.Println("x should have been 0")
  } else if t2 := oneMore(x); t2 == 4 {
    fmt.Println("x should have been 3")
  } else {
    fmt.Println("Not handled")
  }
}

func switchCheck(x int) {
  switch t := oneMore(x); t {
  case 1:
    fmt.Println("x should have been 0")
  case 6:
    fmt.Println("x should have been 5")
  default:
    fmt.Println("Unhandled")
  }
}

func main() {
  ifCheck(0)
  ifCheck(3)
  ifCheck(9)

  switchCheck(0)
  switchCheck(5)
  switchCheck(22)
}
