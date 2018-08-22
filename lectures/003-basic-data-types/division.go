/*
  1) Integer divide: we will find number of students per classroom,
     when we know the total number of students and number of classrooms

  2) Float divide: we will find cost of food per person,
     when we know total cost incurred and the number of persons

  3) x /= 4 is the same as x = x / 4
*/

package main

import "fmt"

func main() {
  var totalStudents, numClassRooms int32 = 225, 15
  var studentsPerClass int32 = totalStudents / numClassRooms
  fmt.Println("Number of students per classroom = ", studentsPerClass)

  var totalCost, numPersons float32 = 397.75, 25.0
  var costPerPerson float32 = totalCost / numPersons
  fmt.Println("Cost of food per person = ", costPerPerson)

  var x int32 = 20
  x = x / 5
  fmt.Println("x = x / 5 is = ", x)

  x = 20 // reset x to 20
  x /= 5
  fmt.Println("x /= 5 is = ", x)
}





















//
