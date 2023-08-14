package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 5; i++ {
				c <- i
			}
			done <- 1
		}()
	}

	go func() {
		for i := 0; i < 10; i++ {
			<-done
		}
		close(c)
	}()

	// for i := 0; i < 50; i++ {
	// 	fmt.Println(<-c)
	// }

	for n := range c {
		fmt.Println(n)
	}

}
