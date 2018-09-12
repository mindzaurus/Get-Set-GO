package main

import (
  "fmt"
)

type AnimalDetails struct {
  weight int // in kgs
  height int // in metres
}

type FoodAndHabitat struct {
  onWater bool // can live in water ?
  onLand  bool // can live on Land?
}

type Elephant struct {
  AnimalDetails  // embedded struct AnimalDetails inside Elephant
  FoodAndHabitat // embedded struct FoodAndHabitat inside Elephant
  origin string  // "African" or "Asian" this field belongs to Elephant struct only
}

type Crocodile struct {
  AnimalDetails  // embedded struct AnimalDetails inside Crocodile
  FoodAndHabitat // embedded struct FoodAndHabitat inside Crocodile
  crocType string // Fresh water, Salt water etc ...
}

func createElephantInfo() {
  var e1 Elephant = Elephant{ }
  e1.weight  = 4000   // => weight comes from AnimalDetails
  e1.height  = 2      // => height comes from AnimalDetails
  e1.onWater = false  // => comes from FoodAndHabitat
  e1.onLand  = true   // => comes from FoodAndHabitat
  e1.origin  = "Asian"
  fmt.Println(e1)

  var e2 Elephant = Elephant { AnimalDetails{5000, 3}, FoodAndHabitat{false, true}, "African" }
  fmt.Println(e2)
}

func createCrocodileInfo() {
  c1 := Crocodile{ AnimalDetails{1000, 7}, FoodAndHabitat{true, true}, "Salt Water" }
  fmt.Println(c1)
}

func main() {
  createElephantInfo()
  createCrocodileInfo()
}
