package main

import "fmt"

// This example here is without the concurrency pipeline
// Using same code in the next file to show concurrency pipeline

func main() {
	c := gen(1, 2, 3, 4)
	out := sq(c)

	for n := range out {
		fmt.Println(n)
	}
	//fmt.Println(<-out)

}

func gen(nums ...int) <-chan int {
	fmt.Println(nums)
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

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

	return out

}
