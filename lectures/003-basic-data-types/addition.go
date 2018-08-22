package main

import "fmt"

func main() {
  var numBoys  int32 = 15
  var numGirls int32 = 16

  // integer addition
  var totalStudents = numBoys + numGirls // note the '+' add and '=' assignment operators
  fmt.Println("Total number of students =", totalStudents)

  // float addition
  hatCost, umbrellaCost := 9.10, 11.20 // Compiler understands float type automatically
  totalCost := hatCost + umbrellaCost  // note the '+' and ':=' operators
  fmt.Println("Total cost =", totalCost)

  // Increment operators x = x + y and x+=
  var x int8 = 5 // Declare and set value of x to 5
  x = x + 5 // Self Increment x
  fmt.Println("x = x + 5 is = ", x)

  x = 0  // Reset x to 0
  x += 5 // Self Increment x, note the "+=" operator
  fmt.Println("x += 5 is = ", x)
}





//
