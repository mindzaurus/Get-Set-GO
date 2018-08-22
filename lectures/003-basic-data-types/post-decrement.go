package main

import "fmt"

func main() {
  var j int32 = 6
  var k int32 = 0

  j--
  k = j

  fmt.Println(j, k)
}

// Only for C, C++, java programmers
// k = j-- // Can't do things like k = j-- or A[k--] etc ....
// --j GO doesn't support pre decrement --i
