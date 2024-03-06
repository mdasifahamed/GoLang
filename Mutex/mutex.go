package main

import (
	"fmt"
	"sync"
)

// `Mutex` stands for  mutual exclusion lock
// Which is used to lock lock memory space for
// certain process to complete it can be read or write
// note: it highly expected that after the lock it should unlock
// otherwise  other procees will be for lack of memory
// it makes the application more faster for  efficient use of the memory by locking the
// memory space for a specifc operation

func main() {
	waitGroup := sync.WaitGroup{}
	mutex := sync.Mutex{}

	fmt.Println("Goroutin is started")

	// slice of the numbers
	numbers := []int{2, 5, 6}
	// slice of the names
	names := []string{"kolin"}

	fmt.Println("Slice of numbers Before Adding From Routine", numbers)
	fmt.Println("Slice of names Before Adding From Routine", names)

	waitGroup.Add(2)
	go func(wg *sync.WaitGroup, mt *sync.Mutex, number ...int) {
		defer wg.Done()
		mt.Lock() // if we dont't use mutex lock here we will get race warning
		numbers = append(numbers, number...)
		mt.Unlock() // it necessaray to unlock memory space after using otherwise other operation might failed due to lack of memory as it was locked
	}(&waitGroup, &mutex, 1, 2, 65, 9, 89, 46, 41, 116, 666, 16, 616, 48, 69, 19, 4, 61, 91, 916, 6161, 6991, 3)

	go func(wg *sync.WaitGroup, mt *sync.Mutex, name ...string) {
		wg.Done()
		mt.Lock()
		names = append(names, name...)
		mt.Unlock()
	}(&waitGroup, &mutex, "jonh", "doe", "martin", "fabiano", "hassel", "ruth", "jain")

	waitGroup.Wait()

	fmt.Println("Slice of numbers After Adding From Routine", numbers)
	fmt.Println("Slice of names After Adding From Routine", names)

}
