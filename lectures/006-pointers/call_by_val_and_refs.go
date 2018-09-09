package main

import (
  "fmt"
)

func addCallByVal(a int, b int) int {
  return a + b
}

func addCallByRef(p *int, q *int) int {
  return *p + *q
}

func adderDemo() {
  var v1 int = 39
  var v2 int = 492

  r := addCallByVal(v1, v2)
  fmt.Println("r := addCallByVal(v1, v2) => ", r)

  r = addCallByRef(&v1, &v2)
  fmt.Println("r = addCallByRef(&v1, &v2) => ", r)
}

func incrementVal(x int) {
  x += 1
  fmt.Println("incrementVal: x is incremented to => ", x)
}

func incrementValRef(x *int) {
  *x += 1
  fmt.Println("incrementVal: x is incremented to => ", *x)
}

func main() {
  adderDemo()

  var k int = 93
  fmt.Println("before incrementVal(k) k is => ", k)
  incrementVal(k)
  fmt.Println("After incrementVal(k) k is => ", k)
  fmt.Println("")

  fmt.Println("before incrementValRef(&k) k is => ", k)
  incrementValRef(&k)
  fmt.Println("After incrementValRef(&k) k is => ", k)
}











//
