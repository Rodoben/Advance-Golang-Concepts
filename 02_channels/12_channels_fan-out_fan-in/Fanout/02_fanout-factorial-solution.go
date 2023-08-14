package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func factorial(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range c {
			out <- fact(n)
		}
		close(out)
	}()

	return out
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

//Fanin Functionmerge takes multiple channels in... 
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(len(cs))
	for _, n := range cs {

		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(n)

	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	n := []int{3, 4, 5, 6, 7, 8, 9, 10}
	c := gen(n...)
	// Fan Out Process happening here!
	f1 := factorial(c)
	f2 := factorial(c)
	f3 := factorial(c)
	f4 := factorial(c)
	f5 := factorial(c)
	f6 := factorial(c)
	f7 := factorial(c)
	f8 := factorial(c)
	f9 := factorial(c)
	f10 := factorial(c)
	var y int
	for n := range merge(f1, f2, f3, f4, f5, f6, f7, f8, f9, f10) {
		y++
		fmt.Println(y, "\t", n)
	}

}
