// Demonstrate use of integers in go

package main

import "fmt"

func main() {
  var numViolinStudents int = 23
  var numFrenchStudents int = 15

  var totalStudents int = numViolinStudents + numFrenchStudents

  fmt.Println("The Total number of Students", totalStudents)
}
