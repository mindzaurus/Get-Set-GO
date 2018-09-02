package main

import "fmt"

func codeBlocksDemo() {
  var i int = 20
  fmt.Println(i)
  { // 1st inner code block
    var i float32 = 56.35
    fmt.Println(i)
    { // 2nd inner code block
      var i string = "golang"
      fmt.Println(i)
    } // end of 2nd inner code block
  } // end of 1st inner code block
}

func main() {
  codeBlocksDemo()
}
