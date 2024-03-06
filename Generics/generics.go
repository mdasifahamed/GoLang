package main

type Ordered interface {
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

// in those generics comes in handy we can define a fucntion with type of any and use it
// as per our required data type

func genericsAdd[T Ordered](a, b T) T {
	return a + b
}

func main() {

}
