
package main

import "fmt"

type Elephant struct { // similar to class
  age int
  weight float32
}

func (e Elephant) print() {
  fmt.Println("Elephant")
  fmt.Println("age: " + fmt.Sprint(e.age))
  fmt.Println("weight: " + fmt.Sprint(e.weight))
}

func main() {
  jumbo := Elephant { age: 20, weight: 2000.5 }
  jumbo.print() // object.method
}
