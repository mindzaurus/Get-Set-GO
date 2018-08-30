package main

import "fmt"

// We will create a function
// which returns a function which
// returns a function
func CounterMakerCreator (x int) ( func() func() int ) { // function which creates new CounterMaker
  return func () ( func() int ) { // Similar to CounterMaker except where was acc := 0
    acc := x
    return func() int { // Counter function
      acc += 1
      return acc
    } //innermost enclosure
  } // 2nd level enclosure
} // 3rd level enclosure

func main() {
  cmc1 := CounterMakerCreator(3) // counter starts at 3 not 0
  c1 := cmc1() // grab the counter from counter maker cmc1
  fmt.Println(c1())
  fmt.Println(c1())
  fmt.Println(c1())
  fmt.Println(c1());  fmt.Println()

  cmc2 := CounterMakerCreator(500) // counter starts at 500 not 0
  c2 := cmc2() // grab the counter from counter maker cmc2
  fmt.Println(c2())
  fmt.Println(c2())
  fmt.Println(c2()) ;  fmt.Println()

  fmt.Println(c1());  fmt.Println()
  fmt.Println(c2()) ;  fmt.Println()

}
