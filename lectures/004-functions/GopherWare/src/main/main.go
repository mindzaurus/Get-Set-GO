
package main

import "fmt"
import "ezmath"
import "ezcfg"
//  Use -> import _ "ezcfg" for a blank import

func main() {
  r := ezmath.AddInt(4, 20)
  fmt.Println("ezmath.AddInt(4, 20)", r)

  r = ezmath.SubInt(24, 19)
  fmt.Println("ezmath.SubInt(24, 19)", r)

  fmt.Println("ezcfg.ClientPort", ezcfg.ClientPort)
}
