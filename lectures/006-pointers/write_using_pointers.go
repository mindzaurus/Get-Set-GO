
package main

import (
  "fmt"
)

var theValue int = 1005

func main() {

  fmt.Println("theValue", theValue)

  var valPtr *int = &theValue // get the address of theValue's memory location
  fmt.Println("theValue is stored in memory location valPtr", valPtr)
  fmt.Println("memory location ",valPtr, "contains the value ",*valPtr)

  *valPtr = 23975 // write value 23975 directly to the memory location
  fmt.Println("memory location ",valPtr, "now has ",*valPtr)

  *valPtr += 3 // increment the value in memory location by 3
  fmt.Println("memory location ",valPtr, "now has incremented value ",*valPtr)
  fmt.Println("Changing memory location ", valPtr, " changed theValue ", theValue)

  // *valPtr++ is the same as *valPtr += 1
  *valPtr++ // using ++ increment operator and increment value in memory location
  fmt.Println("memory location ",valPtr, "has new value ",*valPtr, "bcos of ++ operator")
}








//
