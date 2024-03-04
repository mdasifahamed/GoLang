package main

import "fmt"

func main() {

	// Slices Are Similar To Array But It Has Dynamic Size in Length
	// If The Slice is Created From The Array The Can be Dynamic

	// slice direct declartion
	name_slice := []string{"Sam", "gig", "vis", "team"} // this is similar to array we can define aray like this also
	fmt.Println("Slice Similar to  array", name_slice)

	// array
	numarray := [8]int{2, 98, 5, 9, 6, 7, 9, 7}

	// slice from array

	numSlice := numarray[2:6] // the slice should be [5,9,6,7]

	fmt.Println(numSlice)

	// length of the numSlice should be four

	fmt.Println("Length Of The NumSlcie Is : ", len(numSlice))

	// but has capacity of 6 means that it can can two more elemnets to it
	// as it has dervied from the numarray which has length of 8
	// but we have take the elements form the 2 so it has capacity of 8 - 2 = 6
	fmt.Println("Capasity Of The NumSlcie Is : ", cap(numSlice))

	// We Can Also Modify AnY element in the Slice

	// elments in numarray[3]
	fmt.Println("Before Modifying Slice", numSlice[3])

	// modify the slice with the same index

	numSlice[3] = 100

	fmt.Println("After Modifying Slice", numSlice[3])

	// We Can also Add Elemnets To Slice Using append()

	numSlice = append(numSlice, 200)

	fmt.Println("new numSlice After Adding 1 New Elements", numSlice)
	fmt.Println("new numSlice len After Adding 1 New Elements", len(numSlice))

	numSlice = append(numSlice, 300)

	fmt.Println("new numSlice After Adding Another 1 New Elements", numSlice)
	fmt.Println("new numSlice len After Adding  Another 1 New Elements", len(numSlice))

	// Creating slice using make() it takes three pararmeter dataType, len, and capacity

	ages := make([]int, 2, 4)

	ages[0] = 10
	ages[1] = 20

	fmt.Println("ages", ages)
	// Even We can Append it
	ages = append(ages, 30, 50, 60, 80)
	fmt.Println("ages", ages)

	// We Cannot Append A Array
	// arr1 := [2]int{1, 2}

	// We Cannot Append A Array
	// arr1 = append(arr1, 5) this line will give an error

}
