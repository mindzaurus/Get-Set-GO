package main

import "fmt"

func e() {
  fmt.Println("Reached e")
}

func d() {
  fmt.Println("d calling e")
  e()
  fmt.Println("Returned to d from e")
}

func c() {
  fmt.Println("c calling d")
  d()
  fmt.Println("Returned to c from d")
}

func b() {
  fmt.Println("b calling c")
  c()
  fmt.Println("Returned to b from c")
}

func a() {
  fmt.Println("a calling b")
  b()
  fmt.Println("Returned to a from b")
}

func main() {
  fmt.Println("main calling a")
  a()
  fmt.Println("Returned to main from a")
}
