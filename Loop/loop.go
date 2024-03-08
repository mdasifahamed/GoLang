package main

import "fmt"

func main() {

	// find prime
	numbers := []int{30, 56, 98, 565, 995, 6, 2}

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 100/3 {
			fmt.Println("passed", numbers[i])
		} else {
			fmt.Println("failed", numbers[i])

		}
	}

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 != 0 {
			break
		} else {
			fmt.Println("Continued")
		}

	}
}
