package main

import (
  "fmt"
)

func reWrite(m map[string] string) {
  m["Burrito"] = "Tex Mex"
}

func main() {
  ftMap := map[string] string { }
  ftMap["Burrito"] = "Mexican"  // key => "Burrito", value => "Mexican"

  fmt.Println("ftMap[\"Burrito\"]", ftMap["Burrito"])

  unknown := ftMap["Paella"] // ftMap["Paella"] has no value
  fmt.Println("len(unknown)", len(unknown)) // unknown is string

  reWrite(ftMap) // ftMap passed as reference even when ftMap is not a pointer
  fmt.Println("ftMap[\"Burrito\"]", ftMap["Burrito"])

  ftm2 := map[string] string { // ftm2 -> foodTypeMap2
    "Burrito":      "Mexican", // key: value
    "Paella":       "Spanish", // key: value
    "Kimchi":       "Korean", // key: value
    "Bunny Chow":  "South African", // key: value
    "Sushi":        "Japanese", // key: value
  }

  fmt.Println(ftm2)
  reWrite(ftm2)
  fmt.Println(ftm2)

  fmt.Println("ftm2[\"Paella\"]",ftm2["Paella"])
  fmt.Println("ftm2[\"Bunny Chow\"]",ftm2["Bunny Chow"])
}







//
