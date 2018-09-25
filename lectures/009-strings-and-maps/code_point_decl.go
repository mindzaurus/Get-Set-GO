package main

import (
  "fmt"
)

func main() {
  catStr := "고양이"

  for _, c := range catStr {
    fmt.Printf("%X => %d, ", c, c)
  }
  fmt.Println()

  // ACE0, C591, C774
  s1 := "\uACE0\uC591\uC774" // 16 bit 4 byte code points
  fmt.Println("s1 =", s1)

  // ACE0, C591, C774
  s2 := "\U0000ACE0\U0000C591\U0000C774" // 32 bit 8 byte code points // we will get an error now
  fmt.Println("s2 =", s2)

  s3 := "\uACE0\U0000C591\U0000C774"
  fmt.Println("s3 =", s3)
}











//
