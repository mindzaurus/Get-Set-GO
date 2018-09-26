package main

import (
  "fmt"
  "strings"
)

func main() {
  s := "collège in Zürich"
  fmt.Println(s, ",len(s)", len(s))

  fmt.Println("s[0:7]", s[ 0 : 7 ])
  fmt.Println("s[0:8]", s[ 0 : 8 ])

  fmt.Println("s[12 : 19]", s[12 : 19])

  // Better way to extract and collège Zürich

  // collège is between beginning of string and ends at a space character
  // We will find what is the index of the first space and slice out collège
  SPACE := " "
  i := strings.Index(s, SPACE)
  fmt.Println("s[12 : i]", s[0 : i])

  // Zürich is between last space character and end of string
  // We will find what is the index of the last space and slice out Zürich
  i = 0
  s = "The few of the many collège in Zürich"
  t := s[:] // t is some tmp string

  for {
    i = strings.Index(t, SPACE)
    t = t[ i+1 : ] // slice from next char to space till end of string
    if i < 0 {  // when a space is not found we get a -1 ret value
      break
    }
    fmt.Println(i, t)
  }

}





//
