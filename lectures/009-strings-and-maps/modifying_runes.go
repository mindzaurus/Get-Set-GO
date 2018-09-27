package main

import (
  "fmt"
)

func main() {
  s :=  "collège in Zürich"
  runeSliceOfs := []rune(s)

  for _, ru := range runeSliceOfs {
    fmt.Printf("%q\n", ru)
  }
  fmt.Println("")

  for i := range runeSliceOfs {
    if string(runeSliceOfs[i]) == " " { // don't modify spaces
      continue
    }
    runeSliceOfs[i]-- // modify the rune transpose by one lesser unicode point
  }

  oneLess := string(runeSliceOfs) // create string from modified rune slice
  fmt.Println(oneLess)

  fmt.Println(s)
}






//
