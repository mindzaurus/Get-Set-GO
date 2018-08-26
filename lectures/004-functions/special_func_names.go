// We will use some special symbols like  β , δ  in our program

package main

import "fmt"

func β (x float32) float32 {
  return x * 0.2
}

func δ (x float32) float32 {
  return x * 0.32
}

func main() {
  a := β(5.0)
  fmt.Println("β(5.0)", a)

  b := δ(5.0)
  fmt.Println("δ(5.0)", b)
}
