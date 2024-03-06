package main

import (
	"fmt"
	"sync"
)

// Channels in Go is way to commnuicate and receive data between goroutine
// we can not return data in a regular variable from goroutine
// channels helps us to recevie and send data between go routine

func main() {

	// we create channels using make(chan dataType)

	channel := make(chan int)
	waitGroup := sync.WaitGroup{}
	mutex := sync.Mutex{}
	// we can send through channel using goroutine

	// creating an anoymoun gorouting funtion that sets value to channel
	// sending value to a channels `<-` send arrow
	waitGroup.Add(2)
	go func(wg *sync.WaitGroup, mt *sync.Mutex, ch chan int, nums ...int) {
		defer wg.Done()
		mt.Lock()
		sum := 0
		for _, num := range nums {
			sum += num
		}
		ch <- sum // `<-` is send arrow to channle in channle which mena it sends value channels
		mutex.Unlock()

	}(&waitGroup, &mutex, channel, 1, 8, 9, 7, 98, 7, 9, 6, 56, 16)

	// another anoymous function to recieve data from channels

	go func(wg *sync.WaitGroup, ch chan int) {
		defer wg.Done()

		value := <-ch // this arrow sign is used to receive data from channels

		fmt.Println("Value form the Channel", value)

	}(&waitGroup, channel)

	waitGroup.Wait()

}
