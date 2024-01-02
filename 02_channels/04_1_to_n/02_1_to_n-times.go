package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 500; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for n := range c {
				fmt.Println("i:", i, n)
			}
			done <- true
		}()
	}

	for i := 0; i < 100; i++ {
		<-done
	}
}
