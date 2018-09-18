package main

import (
  "fmt"
)

func printAverage(mSet [5]int) {
  sum := 0.0

  for i := 0; i < len(mSet) ; i++ {
    m := float64(mSet[i])
    sum += m
  }
  fmt.Println("Average ", sum/5.0)
}

func averageRanged(mSet [5]int) {
  sum := 0.0

  for i := range mSet {
    m := float64(mSet[i])
    sum += m
  }

  fmt.Println("Average ranged ", sum/5.0)
}

func avgValueRanged(mSet [5]int) {
  sum := 0.0
  x := 0

  for i, v := range mSet {
    m := float64(v)
    sum += m
    x = i // last iteration will set x = 4
  }
  x++ // number of elements will be last valid index of mSet + 1
  fmt.Println("Average valued range of", x, "items", sum/5.0)
}

func main() {
  markSet1 := [5]int { 100, 96, 40, 68, 89 }
  printAverage(markSet1)

  markSet2 := [5]int { 85, 33, 76, 99, 20 }
  printAverage(markSet2)

  averageRanged(markSet1)
  averageRanged(markSet2)

  avgValueRanged(markSet1)
  avgValueRanged(markSet2)
}
