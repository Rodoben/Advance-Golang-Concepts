package main

import "fmt"

// do factorial for 100 number concurrently

func main() {

	c := gen(3, 4, 5, 6, 7, 8, 9)
	f := factorial(c)

	for n := range f {
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
	f := 1
	for i := n; i > 0; i-- {
		f = f * i
	}
	return f
}
