package main

import "fmt"

// functionRunner takes 1) a function and 2) int value as arguments
func functionRunner ( f func(x int), value int) {
  fmt.Println("I will run a function now says:: functionRunner ")
  f(value)
}


func main() {
  fmt.Println (  "Negative value of 34 is = ",
    func (x int) int {
      return -1 * x
    }(34) )

  // Square then make value negative
  sqThenNeg := func (x int) int { return x * x * -1 }
  fmt.Println ( "Squared then negative value of 46 = ", sqThenNeg(46) )

  functionRunner(func (x int) { fmt.Println("x*x*x", x*x*x) }, 345 )
  functionRunner(func (x int) { fmt.Println("x*x*x*x*x", x*x*x*x*x) }, 345 )
}





//
