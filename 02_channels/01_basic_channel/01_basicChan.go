package main

import (
	"fmt"
	"time"
)

// Channels:
// consider a channel to be a water flowing pipe, once a channel is looped the channel becomes empty

func main() {
	fmt.Println("Channels in action...")
	c := make(chan int)
	var a chan string
	//ch := chan bool //Wrong Way of initializing a channel.
	fmt.Println("Using Make:", c) //address
	fmt.Println("Using a var", a) // nil
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	fmt.Println(c) // address
	// consider a channel to be a water flowing pipe, once a channel is looped the channel becomes empty
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-c)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-c)
	// }

	fmt.Println("______________________________")
	// if we want to make use of range keyword then we have to close the channel as in line no 22
	// for n := range c {
	// 	fmt.Println("From Range:", n)
	// }

	// now if we make use of two go routines to range over the channel, which ever go routine picks the value the result is printed.

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("From 1:", <-c)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("From 2:", <-c)
		}
	}()

	time.Sleep(2 * time.Second)

}
