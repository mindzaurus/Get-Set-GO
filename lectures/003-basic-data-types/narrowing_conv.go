package main

import "fmt"
import "math"

// MAX values
var maxInt16 int16 = math.MaxInt16
var maxInt32 int32 = math.MaxInt32
var maxInt64 int64 = math.MaxInt64
var maxUint32 uint32 = math.MaxUint32
var maxUint64 uint64 = math.MaxUint64
var maxFloat32 float32 = math.MaxFloat32
var maxFloat64 float64 = math.MaxFloat64

// Temporary variables
var i16 int16
var i32 int32
var i64 int64
var ui32 uint32
var ui64 uint64
var f32 float32
var f64 float64

func main() {
  // int32 to int16
  i16 = int16(maxInt32)
  fmt.Println("max of int32 to int16", i16)

  // int64 to int32
  i32 = int32(maxInt64)
  fmt.Println("max of int64 to int32", i32)

  // uint32 to int32
  i32 = int32(maxUint32)
  fmt.Println("max of uint32 to int32", i32)

  // uint64 to uint32
  ui32 = uint32(maxUint64)
  fmt.Println("max of uint64 to uint32", ui32)

  // float64 to float32
  f32 = float32(maxFloat64)
  fmt.Println("max of float64 to float32", f32)
}













//
