package main

import (
  "fmt"
)

func main() {
  var a [2]int
  a[0] = 4
  a[1] = 5
  fmt.Println(a)
  fmt.Printf("type of => var a [2]int = %T\n\n", a)

  var b = [2]int { 6, 7 }
  fmt.Println(b)
  fmt.Printf("type of => var b = [2]int {6, 7} = %T\n\n", b)

  var c = [...] int{ 6, 9, 11 }
  fmt.Println(c)
  fmt.Printf("type of => var c = [...] int{6, 9, 11} = %T\n\n", c)

  var d [2]int = [2]int { -100, -200 }
  fmt.Println(d)
  fmt.Printf("type of => var d [2]int = [2]int { -100, -200 } = %T\n\n", d)

  e := [2]int{ 8,9 }
  fmt.Println(e)
  fmt.Printf("type of => e := [2]int{ 8,9 } = %T\n\n", e)

  f := [3]int { 2: -49, 0: 32, 1: 125 }
  fmt.Println(f)
  fmt.Printf("type of => f := [3]int { 2: -49, 0: 32, 1: 125 } = %T\n\n", f)
}


















//
