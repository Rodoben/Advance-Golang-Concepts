package main

import (
	"fmt"
	"time"
)

func main() {

	// c := make(chan int)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		c <- i
	// 	}
	// 	close(c)
	// }()

	// for n := range c {
	// 	fmt.Println(n)
	// }

	// We cannot assign a chan with var keyword and try all possible ways to range or loop from that channel.
	// <nil> zero value
	//ch := make(chan int)
	var ch chan int
	go func() {
		fmt.Println("i am waiting for a number to be inserted")
		for i := 0; i < 10; i++ {
			fmt.Println("i:", i)
			ch <- i
			fmt.Println(ch)
		}
		//close(ch)
	}()
	// for n := range ch {
	// 	fmt.Println(n)
	// }
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch)
	// }
	time.Sleep(10 * time.Second)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
