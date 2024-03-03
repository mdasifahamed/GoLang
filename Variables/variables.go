package main

import "fmt"

// assigning and initilizing variable at the same time using ':=' this method only works in inside functions
// it cannot defined empty
// it must be defined and initilzed at the same time

// varibale using ':='
func add(a int, b int) int {
	result := a + b
	return result
}

func main() {

	// Varibables in Go

	// named variable with type declration
	// note: go has four has four data types int,float32, string, bool

	// we can use var keyword to define a variable

	var name string = "joe"
	var age int = 25
	var height = 155.5
	var blind = false

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(blind)

	// varibale can be defined without types using var keyword and the compile it will be determiined the type by its value
	var number = "two"
	fmt.Println(number)

	// empty varibles can be defined also but type must be defined

	var b int // defined first
	b = 10    // initialized later
	fmt.Println(b)

	var r = add(10, 20)
	fmt.Println("Addint of two Number", r)

	// using constatnt keyword we can defined constant variable which valu cannot be changed

	const gender string = "male"

	// if we try to
	// gender = "female" we will get compiler error `compilerUnassignableOperand`

	fmt.Println("Constant Varibale", gender)

}
