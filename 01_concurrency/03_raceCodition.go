package main

import (
	"fmt"
	"sync"
)

// Data Race:
// Data race is a condition where one global varibale is taken and we try to increment the counter variable using two go routines.
// Thus go routine read and write to a memory address of the shared counter variable
// Waitgroup helps to syncronize the go routines by adding, waiting and done.

var counter int
var wg sync.WaitGroup

func main() {

	fmt.Println("Data Race case starting.....")
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			for i := 0; i < 10; i++ {

				x := counter
				x++
				counter = x

			}
			wg.Done()
		}()

	}
	//time.Sleep(1 * time.Millisecond)
	wg.Wait()
	fmt.Println("Final Counter:", counter)

}
