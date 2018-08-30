package main

import "fmt"

func A() {

  defer func () { // cleanup function
    fmt.Println("Finished! cleanup running function A()")
  } () // defer needs a function call to b made

  fmt.Println("Executing! function A()")
} // cleanup function executed at this point

func main() {
  A()
}

//
