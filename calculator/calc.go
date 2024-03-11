package main

import (
	"fmt"
	"os"
)

func do_operation() {
	var (
		a float32
		b float32
		c string
	)

	fmt.Println("Insert A Number: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid Input Type", a)
		os.Exit(2)
	}

	fmt.Println("Insert Another Number: ")
	_, err2 := fmt.Scan(&b)
	if err2 != nil {
		fmt.Println("Invalid Input Type", b)
		os.Exit(2)
	}

	fmt.Println("Insert Operation Type i.e(+,-,*,/)")
	_, err3 := fmt.Scan(&c)
	if err3 != nil {
		fmt.Println("Invalid Input Type", c)
		os.Exit(2)
	}

	switch c {
	case "+":
		fmt.Println("Addintion of Two Number", a+b)
	case "-":
		fmt.Println("Subtraction of Two Number", a-b)
	case "*":
		fmt.Println("Mulplication of Two Number", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Divions of by Zero ", a, b)
			os.Exit(3)
		} else {
			fmt.Println("Divions of Two Number", a/b)
		}
	default:
		println("Invalid Opeartion", c)
		os.Exit(4)
	}
}

func main() {

	do_operation()

}
