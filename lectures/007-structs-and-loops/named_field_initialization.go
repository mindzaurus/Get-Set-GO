package main

import (
  "fmt"
)

type Athelete struct {
  height float32 // height should be lesser than 10.0 feet, so height < 10.0
  weight float32 // weight should be lesser than 500 kgs, so weight < 500.0
}

func atheleteInfo(a Athelete) {
  fmt.Println("height", a.height, "weight", a.weight)
}

func main() {
  var a1 Athelete = Athelete{ 5.8/* height */, 75 /* weight */}
  atheleteInfo(a1)

  a2 := Athelete{ 75 /* weight */, 5.8/* height */ } // weight and height wrong order of arguments
  atheleteInfo(a2)

  a3 := Athelete{ height: 5.8, weight: 75}
  atheleteInfo(a3)

  a4 := Athelete{ weight: 85, height: 6.2} // weight and height order of arguments changed
  atheleteInfo(a4)
}
