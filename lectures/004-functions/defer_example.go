package main

import "fmt"

var x int = 1

func printX() {
  fmt.Println(x)
}

func addOne() int {
  printX()
  x += 1
  return x
}

func addOneDefered() int {
  defer printX() // this will push the printX() call inbetween the return x and closing }
  x += 1
  return x
  // printX() -> defer is like executing printX after return x is done
}

func main() {
  v := addOneDefered()
  fmt.Println("main:: v = ", v)
}
