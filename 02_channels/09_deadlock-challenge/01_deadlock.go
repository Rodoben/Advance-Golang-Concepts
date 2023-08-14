package main

import (
	"fmt"
)

// func main() {
// 	c := make(chan int)
// 	c <- 1
// 	fmt.Println(<-c)
// }

// This results in a deadlock.
// Can you determine why?
// And what would you do to fix it?

// Solution:
// making use of buffered channel

// func main() {
// 	c := make(chan int, 1)
// 	c <- 1
// 	fmt.Println(<-c)
// }

// making use of unbuffered channel and go routine

func main() {
	c := make(chan int)

	go func() {
		c <- 2
	}()
	fmt.Println(<-c)
}
