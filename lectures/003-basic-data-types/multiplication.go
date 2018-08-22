/*
  1) Integer multiplication: total number of candies = candies per box * number of boxes
  2) float multiplication: total cost of phones sold = cost per phone * number of phones sold on christmas eve
  3) x *= 2 is same as x = x * 2
*/

package main

import "fmt"

func main() {
  var numCandiesInBox, numBoxes int32 = 33, 45
  var totalCandies int32 = numCandiesInBox * numBoxes
  fmt.Println("Total number of candies = ", totalCandies)

  var phoneCost, numPhonesSold float32 = 176.45, 25.0
  var totalCost float32 = phoneCost * numPhonesSold
  fmt.Println("Total cost of Phones sold on Christmas Eve = ", totalCost)

  var x int32 = 5
  x = x * 2
  fmt.Println("x = x * 2 is = ", x)

  x = 5 // reset x to 5
  x *= 2
  fmt.Println("x *= 2 is = ", x)
}






















//
