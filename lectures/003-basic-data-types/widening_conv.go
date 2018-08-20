package main

import "fmt"
import "math"

// Max values
var maxInt16 int16 = math.MaxInt16
var maxInt32 int32 = math.MaxInt32
var maxUint32 uint32 = math.MaxUint32
var maxFloat32 float32 = math.MaxFloat32

// Temporary variables
var i32 int32
var i64 int64
var ui32 uint32
var ui64 uint64
var f64 float64

func main() {
  // int16 to int32
  i32 = int32(maxInt16)
  fmt.Println("max of int16 to int32", i32)

  // int32 to int64
  i64 = int64(maxInt32)
  fmt.Println("max of int32 to int64", i64)

  // int32 to uint32
  ui32 = uint32(maxInt32)
  fmt.Println("max of int32 to uint32", ui32)

  // uint32 to uint64
  ui64 = uint64(maxUint32)
  fmt.Println("max of uint32 to uint64", ui64)

  // float32 to float64
  f64 = float64(maxFloat32)
  fmt.Println("max of float32 to float64", f64)
}
















//
