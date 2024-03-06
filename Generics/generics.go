package main

import "fmt"

type AnyType interface {
	int | float32 | ~string
}

/*
	generics in go
	Generics is usefull when we don't ehat type of data we are working with
	it helps us in polymorphism capabitlities we can have one function use it for different type
*/

// funtion for adding two integers is
// below the fucntion is fine but what if want to add
// two floats or concane two the we have to write two more functions
func add(a int, b int) int {
	return a + b
}

// in those cases generics comes in handy we can define a fucntion with type of as per our requirments
// in an interface and use those types in tha sam functon to perofrm operations
// as per our required data type
// [T AnyType] `AnyType` interface has three types in it we use those in th following
// / `genericsAdd` function
func genericsAdd[T AnyType](a, b T) T {
	return a + b
}

func main() {

	r := add(12, 52)
	// r:=add(12.5,52.5) // we cannot pass float to it becasuseit takes only int type as parameter

	// but we can do it in `genreicsAdd` function as it uses generics to deetrmine types
	r2 := genericsAdd[int](10, 12)
	r3 := genericsAdd[float32](10.68, 12.694)
	r4 := genericsAdd[string]("GO", "Lang")

	fmt.Println("Addintion of two integer from `add()` ", r)
	fmt.Println("Addintion of two integer from `generics()` ", r2)
	fmt.Println("Addintion of two floats from `generics()` ", r3)
	fmt.Println("Concatation of two stirng from `generics()` ", r4)

}
