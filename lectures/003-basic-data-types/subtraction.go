/*
    1) Integer subtract: number of adults = total visitors in zoo - number of children
    2) float subtract: dollars in hand = total dollars in hand - cost of coffee
    3) -= is a decrement operator x-=2 is same as x = x - 2
*/

package main

import "fmt"

func main() {
  var totalVisitors, numChildren int32 = 79, 36
  var numAdults = totalVisitors - numChildren
  fmt.Println("Number of Adults = ", numAdults)

  // float32
  var dollarsWeHad, cofeeCost float32 = 16.45, 2.20 // total dollars we had, cost of Coffee
  dollarInHand := dollarsWeHad - cofeeCost
  fmt.Println("Dollars in hand after buying coffee = ", dollarInHand)

  x := 5
  x = x - 2
  fmt.Println("x = x - 2 is = ", x)

  x = 5 // Reset Value
  x -= 2
  fmt.Println("x -= 2 is = ", x)
}






















//
