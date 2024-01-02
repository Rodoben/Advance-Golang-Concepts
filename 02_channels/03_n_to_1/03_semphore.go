package main

import "fmt"

// signalling the program
// avoid package sync
// we are again trying to send values to one channel via 2 go routine. and we are not closing the channel in those 2 go routines, so to close the channel we make use of new go routine and first bring out the semaphore
// when using two or more go routines to insert in one chabbel, we dont close the channel, we can do semaphore done <- true
// we make use of another go routine to close the channel and signal out the channel
func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// <-done
	// <-done
	// close(c)
	go func() {
		<-done
		<-done
		close(c)
	}()

	// wrong way of signalling
	// <-done
	// <-done
	// close(c)
	// why? after launching 2 go routines, the control flow immediately want to throw out the value true from the done channel, meanwhile other 2 go routines have not finished its execution.

	// putting the semaphore in another go routines launches a new process and it waits

	// to loop it without range clause , first we do not need to close the chan and can avoid signalling
	// for i := 0; i < 20; i++ {
	// 	fmt.Println(<-c)
	// }

	// Here we are trying to range it one time
	for n := range c {
		fmt.Println(n)
	}
}
