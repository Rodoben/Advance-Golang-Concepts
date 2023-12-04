package main

import "fmt"

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func cubeRoot(c chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range c {
			out <- n * n * n
		}
		close(out)
	}()

	return out
}

func main() {
	for n := range cubeRoot(gen(1, 2, 3, 4)) {
		fmt.Println(n)
	}

	for n := range gen(1, 2, 3, 4) {
		fmt.Println(n)
	}
}
