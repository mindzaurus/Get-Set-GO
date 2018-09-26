package main

import (
  "fmt"
)

func main() {
  s := "college in zurich"
  fmt.Println("len(s)", len(s))

  c := s[ 0 : 7 ] // => college
  fmt.Println("c := s[0:7]", c)

  z := s[ 11 : 17 ] // => zurich
  fmt.Println("z := s[11:17]", z)

  fmt.Println("s[0]", s[0])
  fmt.Printf("s[0] = %q\n", s[0])

  // s[0]++ // strings are immutable can't modify in place
}
