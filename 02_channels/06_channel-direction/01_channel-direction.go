package main

import "fmt"

func main() {
	m := MultiplicationTable()
	for sum := range SumPuller(m) {
		fmt.Println(sum)
	}
}

// We generally do not make use of send-only channel because its difficult to range over that channel
// Hence we make use of receive-only channel and bi-directional channel
// we can typecast the bidirectional-channel to send-only or receive-only channel

func MultiplicationTable() <-chan int {
	out := make(chan int)
	go func() {
		var m int = 1
		for i := 1; i <= 10; i++ {
			m = i * 5
			fmt.Println("M", m)
			out <- m
		}
		close(out)
	}()
	return out
}

func SumPuller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
