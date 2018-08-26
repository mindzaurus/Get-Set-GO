
package main

import "fmt"

func twiceThrice(x int) (int, int) {
  fmt.Println("x*2:", x*2, "x*3:", x*3) // This is side effect no value returned
  return x*2, x*3
}

func main() {
  v1, v2 := twiceThrice(31)
  fmt.Println("v1",v1,"v2",v2)

  _, v2 = twiceThrice(32) // We need only thrice the value now, not twice
  v1, _ = twiceThrice(34) // We need only twice the value now, not thrice

  _, t2 := twiceThrice(10)
  fmt.Println(t2)

  t1, _ := twiceThrice(11)
  fmt.Println(t1)

  _, _ = twiceThrice(12) // ignoring all return values, only side-effects desired
}
