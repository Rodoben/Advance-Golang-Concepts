package main

import (
	"fmt"
	"time"
)

// implement basic concurrency to show the working of go routines.
// usage of time.Sleep to show the diffrent output being printed at different time

func main() {

	fmt.Println("Basic Concurrency...")

	go foo()
	go bar()

	time.Sleep(3 * time.Millisecond)

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
}
