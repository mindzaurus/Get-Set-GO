package main

import (
  "fmt"
)

func twoLvlPtrs() { // two level pointers
  var i1 int = 523
  var p2 *int = &i1
  var p1 **int = &p2

  fmt.Println("i1 is at address ", &i1, "and contains value ", i1)

  fmt.Println("p2 holds address", p2)
  fmt.Println("p1 holds address", p1)

  fmt.Println("Dereferencing two levels of pointer p1 gives value", **p1)
}

func thrLvlPtrs() { // three level pointers
  var f1 float64 = 967.4976
  var q3 *float64 = &f1
  var q2 **float64 = &q3
  var q1 ***float64 = &q2

  fmt.Println("f1 is at address", &f1, "and holds value ", f1)

  fmt.Println("q3 holds address", q3)
  fmt.Println("q2 holds address", q2)
  fmt.Println("q1 holds address", q1)

  fmt.Println("Dereferencing three levels of pointer q1 gives us", ***q1)
}

func main() {
  twoLvlPtrs()
  thrLvlPtrs()
}


















//
