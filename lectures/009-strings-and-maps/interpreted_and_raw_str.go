package main

import (
  "fmt"
)

// He said "I saw a Gopher" and smiled
func interpretedStrings() {
  var s1 string = "He said \"I saw a Gopher\" and smiled"
  fmt.Println(s1)
  fmt.Println("")

  s2 := "Africa has got\nlions\nelephants\nhyenas\ncheetahs\n"
  fmt.Println(s2)
}

var s4 = `Africa has got
lions
elephants
hyenas
cheetahs`

func rawStrings() {
  var s1 string = `He said \"I saw a Gopher\" and smiled`
  fmt.Println(s1)
  fmt.Println("")

  s2 := `Africa has got\nlions\nelephants\nhyenas\ncheetahs\n`
  fmt.Println(s2)
  fmt.Println("")

  s3 := `He said "I saw a Gopher" and smiled`
  fmt.Println(s3)
  fmt.Println("")

  fmt.Println(s4)
  fmt.Println("")

}

func main() {
  interpretedStrings()
  rawStrings()
}






//
