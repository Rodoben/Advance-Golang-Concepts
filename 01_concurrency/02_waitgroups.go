package main

import (
	"fmt"
	"sync"
)

// implement basic concurrenct to show the working of go routines.
// usage of time.Sleep to show the diffrent output being printed at different time

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("Basic Concurrency...")

	go foo()
	go bar()

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
