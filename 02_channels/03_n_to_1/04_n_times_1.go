package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				c <- j
			}
			done <- true
		}()
	}

	// for i := 0; i < 50; i++ {
	// 	fmt.Println(<-c)
	// }

	// by using range clause
	// we need to pull out the semaphore
	// we need to close the channel c
	go func() {
		for i := 0; i < 10; i++ {
			<-done
		}
		close(c)
	}()

	// here we are tring to range it one time

	for v := range c {
		fmt.Println(v)
	}

}
