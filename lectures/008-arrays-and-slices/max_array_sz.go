
package main

import (
  "fmt"
)

// Feel free to play with this by changing MAX_INDEX_INT64 and MAX_INDEX_INT32

// 249974444
const MAX_INDEX_INT64 = 249974444 // max index in my computer for 64 bit int array
//var arr [MAX_INDEX_INT64]int64

const MAX_INDEX_INT32 = 499948888 // max index in my computer for 32 bit int array
var arr [MAX_INDEX_INT32]int32

func main() {
  fmt.Println("MAX_INDEX_INT32", MAX_INDEX_INT32)
  fmt.Printf("%T\n", arr)

  arr2Pointer := new([MAX_INDEX_INT64 * 8]int64)
  fmt.Printf("%T\n", arr2Pointer)
}
