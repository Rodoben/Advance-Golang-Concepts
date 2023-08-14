package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	//fmt.Println(<-c)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Loop", <-c)
		}
		done <- true

	}()
	go func() {
		for n := range c {
			fmt.Println("Range", n)
		}
		done <- true
	}()

	<-done
	<-done
}

// Why does this only print zero?
// And what can you do to get it to print all 0 - 9 numbers?
