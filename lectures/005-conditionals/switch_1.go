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
*/

func makeSoundsSw(x int) {
  switch { // C, C++, java folks note there is no break needed for each case
  case x == 1:
    fmt.Println("cat sound, Meow! Meow!")
  case x == 2:
    fmt.Println("dog sound, Woof! Woof!")
  case x == 4:
    fmt.Println("duck sound, Quack! Quack!")
  case x == 8:
    fmt.Println("sheep sound, Bleat! Bleat!")

  case x == 16:
    fallthrough
  case x == 32:
    fallthrough
  case x == -64:
    fmt.Println("human sound, Yahoo! Yippee!")

  default:
    fmt.Println("Unknown animal, emitting silence")
  }
}

func main() {
  makeSoundsSw(1)
  makeSoundsSw(2)
  makeSoundsSw(4)
  makeSoundsSw(8)
  makeSoundsSw(16)
  makeSoundsSw(32)
  makeSoundsSw(-64)
  makeSoundsSw(512) // undefined sound
  makeSoundsSw(1024) // undefined sound
}

func makeSoundsIf(x int) { // Quite ugly if and else if statements, looks cluttered
  if x == 1 {
    fmt.Println("cat sound, Meow! Meow!")
  } else if x == 2 {
    fmt.Println("dog sound, Woof! Woof!")
  } else if x == 4 {
    fmt.Println("duck sound, Quack! Quack!")
  } else if x == 8 {
    fmt.Println("sheep sound, Bleat! Bleat!")
  } else if x == 16 {
    fmt.Println("human sound, Yahoo! Yippee!")
  } else if x == 32 {
    fmt.Println("human sound, Yahoo! Yippee!")
  } else if x == -64 {
    fmt.Println("human sound, Yahoo! Yippee!")
  } else {
    fmt.Println("Unknown animal, emitting silence")
  }

}

func main1() {
  makeSoundsIf(1)
  makeSoundsIf(2)
  makeSoundsIf(4)
  makeSoundsIf(8)
  makeSoundsIf(16)
  makeSoundsIf(32)
  makeSoundsIf(-64)
}
