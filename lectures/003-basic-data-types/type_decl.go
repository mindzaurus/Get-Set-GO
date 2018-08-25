package main

import "fmt"

type Quantity uint64      // Quantity could be huge number and will be positive
type PeopleCount uint64   // PeopleCount ould be huge number and will be positive

type WholeNumber int64    // WholeNumber could be huge number and may be negative
type Signed32bit int32    // Signed32bit will fit into (2^32-1 ) -1  and may be negative

func main() {
  var numPhones Quantity = 5764
  var numScienceStudents PeopleCount = 779

  fmt.Println("numPhones", numPhones)
  fmt.Println("numScienceStudents", numScienceStudents)

  var unity WholeNumber = 1
  var negativeValue Signed32bit = -1

  fmt.Println("unity", unity)
  fmt.Println("negativeValue", negativeValue)

}
