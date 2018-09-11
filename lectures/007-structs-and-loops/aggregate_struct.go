package main

import (
  "fmt"
)

type TomatoInfo struct {
  quantity int
  isRipe   bool
}

type CreamInfo struct {
  quantity float32
  hasEgg   bool
  isLowFat bool
}

type Soup struct { // Aggregate struct
  tomatoRequired TomatoInfo
  creamRequired  CreamInfo
  waterInML      float32
}

func makeSoup1() { // less tomato more creamy soup
  var tomatoToAdd TomatoInfo = TomatoInfo { 1, true } // 1 tomato, isRipe == true
  creamToAdd  := CreamInfo { 40.0, true, true } // 40 ml, hasEgg == true, isLowFat == true
  waterToAdd  := float32(100.0) // 100 ml of water
  soup := Soup{ tomatoToAdd, creamToAdd, waterToAdd }
  fmt.Println(soup)
}

func makeSoup2() { // more tomato less creamy soup, no egg cream
  soup := Soup { TomatoInfo{2, true}, CreamInfo{20.0, false, true}, 100.0 }
  fmt.Println(soup)
}

func main() {
  fmt.Println("less tomato more creamy soup")
  makeSoup1()
  fmt.Println()

  fmt.Println("more tomato less creamy soup, no egg cream")
  makeSoup2()
  fmt.Println()
}
