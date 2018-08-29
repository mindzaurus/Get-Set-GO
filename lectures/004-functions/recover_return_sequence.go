package main

import "fmt"

func eDefer() { fmt.Println("eDefer") } // not called
func e() { // not called
  defer eDefer() // not called
  fmt.Println("Reached e")
}
func dDefer() { fmt.Println("dDefer") } // not called
func d() {  // not called
  defer dDefer() // not called
  fmt.Println("d calling e"); e(); fmt.Println("Returned to d from e")
}

func c1Defer() { fmt.Println("c1Defer") }
func c2Defer() { fmt.Println("c2Defer") }
func c() {
  defer c1Defer()
  defer c2Defer()
  panic("Panicking at c() ")
  fmt.Println("c calling d")
  d()
  fmt.Println("Returned to c from d")
}

func b3Defer() { fmt.Println("b3Defer") }
func b2Defer() { fmt.Println("b2Defer") }
func b1Defer() { fmt.Println("b1Defer") }
func b() {
  defer b1Defer()
  defer b2Defer()
  defer b3Defer()
  fmt.Println("b calling c")
  c()
  fmt.Println("Returned to b from c")
}

func aDefer() { fmt.Println("aDefer");
  recover() /* with recover we return to main normally */
}
func a() {
  defer aDefer()
  fmt.Println("a calling b")
  b()
  fmt.Println("Returned to a from b")
} // aDefer() called here

func mainDefer() { fmt.Println("mainDefer") }
func main() {
  defer mainDefer()
  fmt.Println("main calling a")
  a()
  fmt.Println("Returned to main from a")
}
