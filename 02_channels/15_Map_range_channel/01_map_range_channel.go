package main

import "fmt"

func main() {
	m := map[int]int{
		1:  45,
		5:  56,
		72: 76,
	}

	ch := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, n := range m {
			ch <- n
			ch2 <- 1
		}

	}()

	go func() {
		var i int
		for n := range ch2 {
			i += n
			if i == len(m) {
				close(ch)
			}
		}
	}()

	for n := range ch {
		fmt.Println(n)
	}

	close(ch2)

}
