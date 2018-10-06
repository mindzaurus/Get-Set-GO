package main

import (
  "fmt"
  "time"
  "sync"
)

var x = 0
var m sync.Mutex

func inc() { x++ }
func dec() { x-- }

func doNum(numTimes int, f func()) {
  defer func() { fmt.Println("Done!") } ()

  for j := 0 ; j < numTimes; j++ {
    go func() {
      for i := 0; i < 2; i++ {
        m.Lock()
        f()
        m.Unlock()
      }
    }()
  }
}


func main() {
  nt := 1000 // nt = numtimes
  doNum(nt, inc)
  doNum(nt, dec)

  time.Sleep(100 * time.Millisecond)
  fmt.Println("x = ", x) // correct expected value is 0
}














//
