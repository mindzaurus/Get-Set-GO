package main

import "fmt"
import "math"
// All signed ints int8, int16, int32, int64 have same wrap around behaviour
// All unsigned uints uint8, uint16, uint32, uint64 have same wrap around behaviour
// All floats float32, float64 have same wrap around behaviour -> stays put at max value
func main() {
  var i int32 = math.MaxInt32
  var ui uint32 = math.MaxUint32
  var f float32 = math.MaxFloat32

  fmt.Println("Before incrementing i (int32) ", i)
  fmt.Println("Before incrementing ui (uint32) ", ui)
  fmt.Println("Before incrementing f (float32) ", f, "\n")

  i = i + 1
  ui = ui + 1
  f = f + 1.0

  fmt.Println("After incrementing i (int32) ", i)
  fmt.Println("After incrementing ui (uint32) ", ui)
  fmt.Println("After incrementing f (float32) ", f)
}
