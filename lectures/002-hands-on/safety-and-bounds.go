package main

func incrementInt(v int) int {
  return v + 1
}

func main() {
  triggerError := false
  incrementInt(3)
  //incrementInt(3.3) // warning issued

  if triggerError == true {
    // incrementInt("bad_value")
  }

  k := 4

  ar := [3] int { 3, 2, 5 }
  ar[2] = 20 // OK! accessing last element (3rd element)
  ar[k] = 20 // BAD! accessing 5the element in a 3 element array
}
