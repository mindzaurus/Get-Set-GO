package main

import (
  "fmt"
  "strconv"
)


func main() {
  // func Atoi(s string) (int, error)
  v1, _ := strconv.Atoi("5889")
  fmt.Println(v1+3)

  // func Itoa(i int) string
  v2 := strconv.Itoa(v1*3)
  fmt.Println(v2+"/-")

  // func ParseInt(s string, base int, bitSize int) (i int64, err error)
  v3, _ := strconv.ParseInt("8463729", 10, 64)
  fmt.Println(v3+10)

  v3, _ = strconv.ParseInt("ff83729", 16, 64)
  fmt.Println(v3)

  // func ParseFloat(s string, bitSize int) (float64, error)
  v4, _ := strconv.ParseFloat("328738.8372", 64) // float64
  fmt.Println(v4+100.25)

  v5, _ := strconv.ParseFloat("738136.72783", 32) // float32
  fmt.Println(v5+25.25)
}

















//
