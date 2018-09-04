package main

import "fmt"

type Animal int
const (
  CAT Animal = iota
  DOG
  DUCK
  SHEEP
  COW
)

func makeAnimalSounds(a Animal) {
  switch a {
  case CAT:
    fmt.Println("Meow! Meow!")

  case DOG:
    fmt.Println("Woof! Woof!")

  case DUCK:
    fmt.Println("Quack! Quack!")

  case SHEEP:
    fmt.Println("Bleat! Bleat!")

  case COW:
    fmt.Println("Moo! Moo!")
    
  default:
    fmt.Println("UNKNOWN_ANIMAL silence!")
  }
}

func main() {
  makeAnimalSounds(CAT)
  makeAnimalSounds(DOG)
  makeAnimalSounds(DUCK)
  makeAnimalSounds(SHEEP)
  makeAnimalSounds(COW)
}
