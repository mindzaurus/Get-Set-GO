// Demonstrate  use of basic float data type
// C, Java, C++ programmers not no "float" in go use either float32 or float64

package main

import "fmt"

func main() {
  var cofeePrice float32 = 4.65
  var cookiePrice float32 = 8.95

  var totalPrice float32 = cofeePrice  + cookiePrice

  fmt.Println("Total price to pay ", totalPrice)
}
