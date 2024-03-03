package main

import "fmt"

func main() {

	// Array are fixed length in go
	// it supported single data type
	// one array cannot have multiple data types

	// syntax to declare array is [length]datatype{elments}

	var names = [2]string{"Itel", "Symphony"}

	fmt.Println(names)
	// access array elment by index

	fmt.Println("First Elelment", names[0])
	fmt.Println("Second Elelment", names[1])

	// declaring array without length but it will determined at compile time by the compiler.

	ages := []int{10, 20, 30} // definig and initilizing in a sinlge line without using var keyword using ':='

	fmt.Println(ages)

}
