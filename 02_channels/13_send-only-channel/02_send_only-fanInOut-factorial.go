package main

import (
	"fmt"
	"time"
)

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			for j := 3; j <= 10; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out

}

func fanOut(in <-chan int, n int, out chan<- int) {
	for i := 0; i < n; i++ {
		factorial(in, out)
	}
}

func factorial(in <-chan int, out chan<- int) {
	go func() {
		for n := range in {
			out <- fact(n)
		}
	}()
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {

	in := gen()
	out := make(chan int)

	fanOut(in, 10, out)

	go func() {
		for n := range out {
			fmt.Println(n)
		}

	}()

	time.Sleep(1 * time.Second)
}
