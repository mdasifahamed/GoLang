package main

import (
	"fmt"
	"time"
)

func proposal1(channel1 chan string) {

	time.Sleep(5 * time.Second)
	channel1 <- "proposal 1"
}

func proposal2(channel2 chan string) {

	time.Sleep(6 * time.Second)
	channel2 <- "proposal 1"
}

func proposal3(channel2 chan string) {

	time.Sleep(2 * time.Second)
	channel2 <- "proposal 3"
}

func main() {

	r1 := make(chan string)
	r2 := make(chan string)
	r3 := make(chan string)

	go proposal1(r1)
	go proposal2(r2)
	go proposal3(r3)

	select {
	case op1 := <-r1:
		fmt.Println(op1)

	case op2 := <-r2:
		fmt.Println(op2)

	case op3 := <-r3:
		fmt.Println(op3)

	}

}
