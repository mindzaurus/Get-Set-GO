  package main

  import (
    "fmt"
    "time"
  )

  // To sleep for e.g 5 seconds call waitForSeconds(5)
  func waitForSeconds(seconds int) {time.Sleep(time.Duration(seconds) * time.Second )}

  func runServer() {
    panic("Complete filesystem  unreadable")
  } // just panic

  func shutDown(msg string, caller string) { fmt.Println(msg, caller) }

  func recoveryFunction() {
    var e interface {} = nil // don't bother about this line we will study interface in detailed way
    e = recover() // this is the key line in whole program

    // if e is nil then no panic occured in program
    if e == nil { // we will see if later don't worry about this line
      return
    }

    // panic occured because e is not nil and we will call emergency shutDown
    shutDown("Doing emergency shutdown", "recoveryFunction")
  }

  func main() {
    defer recoveryFunction() // call at the end of main function
    runServer()
    shutDown("Doing graceful shutdown", "main")
    waitForSeconds(3)
  } // recoveryFunction called at this point
