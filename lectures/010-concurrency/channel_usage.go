package main
import (
  "fmt"
)

// primeCh - channel where we will send primes
// doneCh - send a completion message after receiving requested primes
func getPrimes(numPrimes int, primeCh chan int, doneCh chan string) {
  p := []int { 2, 3, 5, 7, 13, 17, 19, 23, 29, 31, 37 }

  for i := 0; i < numPrimes; i++ {
    primeCh <- p[i] // write prime to the primes channel
  }

  doneCh <- "Yeah! done"
}

func main() {
  N := 9 // how many primes we want
  primes := make (chan int) // channel where we will get primes
  done := make(chan string) // get a completion message after receiving requested primes

  go getPrimes(N, primes, done)

  for i := 0; i < N; i++ {
    fmt.Println("prime ", <-primes )// read primes one by one from the primes channel
  }

  // main function waits / blocks till it gets a message in the done channel
  fmt.Println(<-done) // get the completion message
}






//
