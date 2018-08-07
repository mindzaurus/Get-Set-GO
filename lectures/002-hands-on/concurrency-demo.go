package main

import "fmt"
import "time"
import "strconv"

func main () {
  for i := 0; i < 20; i++ {
    go concurrentFunction(i) // This line is the crux of GO
  }
  waitForSeconds(5)
  fmt.Println("\nDone with all concurrent functions")
}

func waitForSeconds (seconds int) {
  waitFor := time.Duration(seconds) * time.Second
  time.Sleep(waitFor)
}

func concurrentFunction (routineIdNumber int) {
  fmt.Println("going to run concurrent function #" +
    strconv.Itoa(routineIdNumber))
}
