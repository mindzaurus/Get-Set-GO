
package main

import (
  "fmt"
)

func main() {

  v1 := 24
  fmt.Printf("v1 = %d\n", v1)
  fmt.Printf("v1 = %d v1+1 = %d\n", v1, v1+1)

  v2 := 9833.3865
  fmt.Printf("v2 = %f\n", v2)

  v3 := "The not so long string"
  fmt.Printf("double quoted v3 = %q\n", v3) // double quoted string
  fmt.Printf("unquoted v3 = %s\n", v3) // raw utf8 bytes

  fmt.Printf("type of \nv1 = %T\nv2 = %T\nv3 = %T\n\n", v1, v2, v3)
  fmt.Printf("type of \n\tv1 = %T\n\tv2 = %T\n\tv3 = %T\n", v1, v2, v3)
}









//
