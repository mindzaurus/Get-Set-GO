package main

import "fmt"

// CounterMaker returns a -> function which returns int
func CounterMaker() func () int {
  acc := 0 // variable declared in the CounterMaker function's environment
  return func () int { // enclosure 1 begin
    acc += 1  // acc accesed inside enclosure 1 which is from CounterMaker's environment
    return acc // acc is trapped inside the environment of enclosure 1
  } // enclosure 1 end
}


func main() {
  c := CounterMaker() // c is now a function holder not a value holder
  fmt.Println( c() )  // c can be called like a normal function like c()
  fmt.Println( c() )
  fmt.Println( c() )
  fmt.Println()

  // d will get a new environment for CounterMaker so acc will be set to 0 for its env
  d := CounterMaker() // d is now a function holder not a value holder
  fmt.Println( d() )  // d can be called like a normal function like d()
  fmt.Println( d() )
  fmt.Println( d() )
  fmt.Println()

  fmt.Println( c() ) // c's environment is different from d, so expect c to retain its old state
}
















//
