package main

import (
  "fmt"
)

func grtLessOrEql(x, y int) {
  if x > y { // check if x is greater than y
    fmt.Println("x",x,"y",y,"x is greater than y")
  } else if x < y { // check if x is lesser than y
    fmt.Println("x",x,"y",y,"x is lesser than y")
  } else { // if x is not greater than or lesser than y -> then x is equal to y
    fmt.Println("x",x,"y",y,"x is equal to y")
  }
}

func main() {
  grtLessOrEql(15, 2)
  grtLessOrEql(3, 17)
  grtLessOrEql(-4, -4)
}
