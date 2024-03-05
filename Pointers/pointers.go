package main

import "fmt"

func main() {
	//A pointer is a special kind of variable that is not only used to store the memory addresses of other variables
	//but also points where the memory is located and provides ways to find out the value stored at that memory location

	// * Operator also termed as the dereferencing operator used to
	// declare pointer variable and access the value stored in the address.

	// & operator termed as address operator
	//used to returns the address of a variable or to access the address of a variable to a pointer.

	// We Use `&` to define a pointer variable
	// We Use `*` operator to pointer variable to access the value stored
	// at the variable at which it is pointing

	var num int = 100

	accessVar := &num // now access var has the memory address of the num variable

	fmt.Println("Memory Address of The num from pointer", accessVar)
	fmt.Println("Memory Address of The num from num", &num)
	// we can update num variable value from the ointer variable accessVar using `*` operator

	// before update of the value of the num

	fmt.Println("Value of the num", num)

	//we can read from the accessVar usig `*`
	fmt.Println("Reading num from  accessVar ", *accessVar)

	// update num from  accessVar using `*`
	*accessVar = 200

	// value after update
	fmt.Println("Value of num after update  ", num)

}
