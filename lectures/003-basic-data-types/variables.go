package main

import "fmt"

func main() {
	var valueA int32			// Only Declaration
	valueA = 10					// Value set for variable
	fmt.Println(valueA)

	var valueB int32 = 20		// Declaration and Initialization with type information
	fmt.Println(valueB)

	var valueC = 30				// Declaration and Initialization without type information
	fmt.Println(valueC)

	var (						// var block multiple variable initialization
		valueD = 40				// No type mentioned
		valueE string = "Tom"   // Type information added "string"
	)
	fmt.Println(valueD, valueE)

	var valueF, valueG int32 = 50, 60 // Multiple variables with type and comma ',' used
	fmt.Println(valueF, valueG)

	var valueH, valueI = 70, "Jerry"  // Multiple different type variables without type info, ',' used
	fmt.Println(valueH, valueI)

	valueJ := 80					  // Short form declaration and initialization
	fmt.Println(valueJ)

	valueK, valueL := 90, "Mickey"
	fmt.Println(valueK, valueL)
}

























//