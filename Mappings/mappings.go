package main

import "fmt"

// Mappings re similar to pyhton dictinary on go
// It stores valu in key value pair
// To define a mpping use map[keyType]vallutype

func main() {

	map1 := map[string]int{"A": 1, "B": 2}

	fmt.Println(map1)

	// Access map value by its key

	fmt.Println("First Key Value Is ", map1["A"])
	fmt.Println("Second Key Value Is ", map1["B"])

}
