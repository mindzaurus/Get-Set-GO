package main

import "fmt"
import "math"

func main() {
  var maxFloat32 float32 = math.MaxFloat32
  var minFloat32 float32 = math.SmallestNonzeroFloat32 // no MinFloat32
  fmt.Println("Maximum value of float32 ", maxFloat32)
  fmt.Println("Minimum value of float32 ", minFloat32, "\n")

  var maxFloat64 float64 = math.MaxFloat64
  var minFloat64 float64 = math.SmallestNonzeroFloat64 // no MinFloat64
  fmt.Println("Maximum value of float64 ", maxFloat64)
  fmt.Println("Minimum value of float64 ", minFloat64, "\n")
}
