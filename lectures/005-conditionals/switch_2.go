package main

import (
  "fmt"
)

/*
  1 -> cat
  2 -> dog
  4 -> duck
  8 -> sheep

  16  -> kumar
  32  -> katie
  -64 -> kim

  horse and pony have same sounds
  128 -> horse
  256 -> pony
*/

func makeAnimalSound(x int) {
  switch x { // note x being used after switch
  case 1: // note the value directly added here
    fmt.Println("Meow!")
  case 2:
    fmt.Println("Woof!")
  case 4:
    fmt.Println("Quack!")
  case 8:
    fmt.Println("Bleat!")

  case 16:
    fallthrough
  case 32:
    fallthrough
  case -64:
    fmt.Println("Yahoo! Yippie")

  // alternate for a fallthrough
  case 128, 256: // note the , this handles both 128 (horse) and 256 (pony)
    fmt.Println("Neigh! Neigh!")

  default:
    fmt.Println("Silence!")
  }
}

func main() {
  makeAnimalSound(1)
  makeAnimalSound(2)
  makeAnimalSound(4)
  makeAnimalSound(8)
  makeAnimalSound(16)
  makeAnimalSound(32)
  makeAnimalSound(-64)

  makeAnimalSound(128)
  makeAnimalSound(256)

  makeAnimalSound(5002)
}
