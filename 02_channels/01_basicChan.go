package main

import (
	"fmt"
	"time"
)

//Channels:

func main() {
	fmt.Println("Channels in action...")
	c := make(chan int)
	var a chan string
	//ch := chan bool //Wrong Way of initializing a channel.
	fmt.Println("Using Make:", c)
	fmt.Println("Using a var", a)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		//close(c)
	}()

	fmt.Println(c)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-c)
	// }
	fmt.Println("______________________________")

	// for n := range c {
	// 	fmt.Println("From Range:", n)
	// }

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
