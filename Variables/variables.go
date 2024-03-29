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

	var num int // defined first
	num = 10    // initialized later
	fmt.Println(num)

	var r = add(10, 20)
	fmt.Println("Addint of two Number", r)

	// using constatnt keyword we can defined constant variable which valu cannot be changed

	const gender string = "male"

	// if we try to
	// gender = "female" we will get compiler error `compilerUnassignableOperand`

	fmt.Println("Constant Varibale", gender)

	// multiple varibale in a sinlge line with out types

	var a, b, c, d, e = 10, 20.5, "Asif", true, 60

	fmt.Println("From multiple varibale a ", a)
	fmt.Println("From multiple varibale b ", b)
	fmt.Println("From multiple varibale c ", c)
	fmt.Println("From multiple varibale d ", d)
	fmt.Println("From multiple varibale e ", e)

	// multiple varible with types and inside blokcs then initilize separately

	var (
		check bool
		total int
	)

	check = false
	total = 50

	fmt.Println("Multiple Varibales From Blocks And Later Initlized", check)
	fmt.Println("Multiple Varibales From Blocks And Later Initlized", total)

}
