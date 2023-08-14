package main

import "fmt"

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
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}

	return total
}

func main() {
	n := []int{3, 4, 5, 6, 7, 8, 9, 10}
	c := gen(n...)
	out := factorial(c)
	for a := range out {
		fmt.Println(a)
	}
}
