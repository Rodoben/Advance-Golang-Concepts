package main

import (
	"fmt"
	"sync"
)

// Data Race:
// Data race is a condition where one global varibale is taken and we try to increment the counter variable using two go routines.
// Thus go routine read and write to a memory address of the shared counter variable
// Waitgroup helps to syncronize the go routines by adding, waiting and done.
// Mutex helps to resolve the data race by locking and unlocking the variable when being manipulated.

var counter int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {

	fmt.Println("Data Race case starting.....")
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	//time.Sleep(1 * time.Millisecond)
	wg.Wait()
	fmt.Println("Final Counter:", counter)

}

func incrementor(s string) {
	for i := 0; i < 10; i++ {
		mu.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mu.Unlock()
	}
	wg.Done()

}
