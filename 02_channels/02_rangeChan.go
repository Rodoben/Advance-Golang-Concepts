package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

	// We cannot assign a chan with var keyword and try all possible ways to range from that channel.
	// <nil> zero value
	var ch chan int
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//close(ch)
	}()
	// for n := range ch {
	// 	fmt.Println(n)
	// }
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println(ch)
}
