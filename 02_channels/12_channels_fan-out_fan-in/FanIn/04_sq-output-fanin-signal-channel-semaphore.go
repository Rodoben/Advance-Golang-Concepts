package main

import (
	"fmt"
)

func main() {

	c1 := gen(2, 3, 4, 5)
	c2 := gen(6, 7, 8, 9)

	s1 := sq(c1)
	s2 := sq(c2)

	for n := range merge(s1, s2) {

		fmt.Println(n)

	}

}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	fmt.Println("I generated numbers")
	return out

}

func sq(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()
	fmt.Println("I generated square")

	return out
}

//Fan In

func merge(c ...<-chan int) <-chan int {
	out := make(chan int)
	done := make(chan bool)

	for _, n := range c {
		go func(ch <-chan int) {
			for n1 := range ch {
				out <- n1
			}
			done <- true

		}(n)
	}
	go func() {
		for range c {
			<-done
		}
		close(out)
	}()

	return out
}
