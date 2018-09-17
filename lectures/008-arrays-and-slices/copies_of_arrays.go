package main

import (
  "fmt"
)

func arrayHandler(arr [3]int) {
  fmt.Println("arrayHandler:: before any change arr=", arr)
  arr[0] = 93
  fmt.Println("arrayHandler:: After arr[0] = 93 arr=", arr)
}

func main() {
  arr01 := [...] int { 11, 22, 33 }
  fmt.Println(arr01)

  arr02 := arr01
  fmt.Println(arr02)

  arr02[1] = 99

  fmt.Println("After setting arr02[1] = 99, arr01=", arr01)
  fmt.Println("After setting arr02[1] = 99, arr02=", arr02)
  fmt.Println("")

  arr03 := [3]int { 101, 203, 305 }
  fmt.Println("Initialized fresh array arr03", arr03)
  arrayHandler(arr03)
  fmt.Println("After arrayHandler(arr03)", arr03)
}
