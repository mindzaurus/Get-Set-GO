package main

import "fmt"
import "math"

// There is no math.MinUint constants as
// the minimmum value for each unsigned integer type is 0 (zero)
func main() {
  // uint8
  var maxUint8 uint8 = math.MaxUint8
  fmt.Println("Maximum value of uint8 ", maxUint8, "\n")

  // uint16
  var maxUint16 uint16 = math.MaxUint16
  fmt.Println("Maximum value of uint16 ", maxUint16, "\n")

  // uint32
  var maxUint32 uint32 = math.MaxUint32
  fmt.Println("Maximum value of uint32 ", maxUint32, "\n")

  //uint64
  var maxUint64 uint64 = math.MaxUint64
  fmt.Println("Maximum value of uint64 ", maxUint64, "\n")
}
