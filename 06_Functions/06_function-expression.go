package main

import "fmt"

func main() {

	func() {
		fmt.Println("I am anonymous func")
	}()
	c := make(chan int, 2)
	f := func(s chan int) {
		for i := 0; i < 2; i++ {
			s <- i
		}
		close(c)
	}
	// f(c)
	go f(c)
	for v := range c {
		fmt.Println(v)
	}

}
