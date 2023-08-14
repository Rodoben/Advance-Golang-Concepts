package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < 1000; i++ {

		go func() {

			for n := range c {
				time.Sleep(1 * time.Millisecond)
				fmt.Println("\t GoRoutine No:", i)
				fmt.Println(n)

			}
			done <- true
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}

	//time.Sleep(1 * time.Second)

}
