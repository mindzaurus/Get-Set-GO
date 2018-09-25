package main

import (
  "fmt"
)

func main() {

  catStr := "고양이"
  catBytes := []byte(catStr)

  for i, b := range catBytes {
    fmt.Printf("index = %d, byte value = %d \n", i, b)
  }

  s1 := []byte{ 234, 179, 160 }
  fmt.Println(string(s1))

  s2 := []byte { 236, 150, 145 }
  fmt.Println(string(s2))

  s3 := []byte { 236, 157, 180 }
  fmt.Println(string(s3))

  s4 := []byte { 234, 179, 160, 236, 150, 145, 236, 157, 180 }
  fmt.Println("string(s4)", string(s4))
}







//









//
