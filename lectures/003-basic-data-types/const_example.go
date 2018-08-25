package main

import "fmt"

func main() {
  const π = 3.14
  const numDaysOfWeek = 7

  fmt.Println("π = ", π)
  fmt.Println("numDaysOfWeek = ", numDaysOfWeek)

  const (
    numJanDays = 31
    numMarDays = 31
    numJunDays = 30
  )

  fmt.Println("Jan has ", numJanDays, "\nMar has ", numMarDays, "\nJun has ", numJunDays)
}
