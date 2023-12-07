package main

import (
	"fmt"
	"sync"
)

// n number of go routines writing to one channel
// using waitgroups in correct way to avoid data race

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// not closing this channel here because we are trying to launch two go routines to insert numbers into one channel
		// we make use of waitgroup.Done() to finish the execution
		wg.Done()
	}()
	// go routine 1 trying to put numbers in same channels as other channel
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 20; i > 0; i-- {
			c <- i
		}
		wg.Done()
	}()

	// when it comes to go routine with waitgroups to wait for go routines, we wait for the exection of operation in fresh go routine and then close the channel when the operation id done
	go func() {
		wg.Wait()
		close(c)
	}()
	// wg.Wait()
	// close(c)
	// time.Sleep(time.Second)

	for n := range c {
		fmt.Println(n)
	}
}
