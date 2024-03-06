package main

//
import (
	"fmt"
	"sync"
)

/*
Goroutine is a routine work that can proccess multiple task
at the same time. To define a proccess as go routine use
`go` keyword before the function to make that function as goroutine
goroutine is managed by go runtime.
*/

/*
`sync` is a proccess managemnt package
it has waitgroup function
which help in running goroutine at the main function.
waitgroup.wait() tells the main to wait untill all the operations
is done inside the main function as goroutine procees takes time to proccess
if we dont use waitGropu we cannnot catch the result of the goroutine process
inside the main
*/
func loop10Times(wg *sync.WaitGroup) {
	// it says that the proccess of this function is done.
	// As statement  after `defer` is executed at the end of the block
	// so the it will be executed ata the end of the function.
	defer wg.Done()
	for i := 0; i < 10000000; i++ {
		if i == 9999999 {
			fmt.Println("I am done running 10000000 times loops")
		}
	}
}

func loop20Times(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 200000000; i++ {
		if i == 199999999 {
			fmt.Println("I am done running 200000000 times loops")
		}
	}
}

func loop30Times(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 300000000; i++ {
		if i == 299999999 {
			fmt.Println("I am done running 300000000 times loops")
		}
	}
}

func main() {
	fmt.Println("Gouroutine  is Started")

	waitGroup := sync.WaitGroup{} // definig wait group

	waitGroup.Add(3) // we have three goroutine so we have to three proccess to the waitgruop
	go loop10Times(&waitGroup)
	go loop30Times(&waitGroup)
	go loop20Times(&waitGroup)

	waitGroup.Wait() // this is the line that says to the  main function to untill the goroutine proccess is finished

	fmt.Println("Gouroutine  is Stopped")

}
