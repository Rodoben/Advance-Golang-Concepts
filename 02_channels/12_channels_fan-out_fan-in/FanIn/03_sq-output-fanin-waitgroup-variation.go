package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen(1, 2)
	in1 := gen(3, 4)
	c1 := sq(in)
	c2 := sq(in1)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
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
	fmt.Println("I generated sqaure")
	return out
}

func merge(ch ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(len(ch))
	/// Just a small variation here is to initaialize a func into a varibale nad then launch it as a goroutine.
	output := func(c <-chan int) {
		for n1 := range c {
			out <- n1
		}
		wg.Done()

	}

	for _, n := range ch {
		go output(n)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
