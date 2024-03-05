package main

import "fmt"

func changeValueUsingPointer(a *int) {

	*a = 100
}

type man struct {
	name  string
	age   int
	blind bool
}

func (m *man) display() {

	fmt.Println("name:", m.name)
	fmt.Println("age:", m.age)
	fmt.Println("blind:", m.blind)

}

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

	// using pointer in the function

	/*
		The `changeValueUsingPointer` takes a pointer varibal changes cgahe value of the given
		variable to the function
	*/

	// defining a variable

	y := 325

	fmt.Println("Value of 'Y' before passing to the function", y)

	changeValueUsingPointer(&y) // passing the address of y to the fucntion now it chnage its value to 100
	fmt.Println("Value of 'Y' after passing to the function", y)

	// we can use pointer inther struct also

	// create a man struct

	man1 := man{name: "joe", age: 25, blind: false}

	// craete a poointer to man struct
	manPoint := &man1

	// now we can use `manPoint` to access and update man1
	// this line shlud print all the info the man struct
	manPoint.display()

	// we can also update it
	manPoint.name = "jordan"
	manPoint.blind = true

	// after updating

	manPoint.display()

}
