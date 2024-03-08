package main

import (
	"asif/greetings"
	"fmt"
)

func main() {

	message := greetings.Hello("GO")

	fmt.Println(message)
}
