package main

import (
  "fmt"
)

type Student struct {
  name string
  age  int
}

func agePlusOne(s Student) Student {
  s.age += 1
  return s
}

func main() {
  s1 := Student {"Tom", 30}
  fmt.Println("main: s1 before agePlusOne ", s1)

  s2 := agePlusOne(s1)
  fmt.Println("main: s2 returned from agePlusOne ", s2)

  fmt.Println("main: s1 after agePlusOne ", s1)
}
