package main

import (
	"fmt"
	"sync"
)

// n number of go routines writing to one channel

var wg sync.WaitGroup

type hotdog struct {
	number int
	name   string
}

func main() {
	c := make(chan hotdog)
	//wg.Add(2)  we need to specify the add method outside the go rotuine to avoid race condition
	xs := []hotdog{
		{number: 1, name: "Goroutine"},
		{number: 2, name: "Goroutine"},
		{number: 3, name: "Goroutine"},
		{number: 4, name: "Goroutine"},
		{number: 5, name: "Goroutine"},
		{number: 6, name: "Goroutine"},
		{number: 7, name: "Goroutine"},
		{number: 8, name: "Goroutine"},
		{number: 9, name: "Goroutine"},
		{number: 10, name: "Goroutine"},
		{number: 11, name: "Goroutine"},
	}
	go func() {
		wg.Add(1)
		for _, x := range xs {
			x.name = x.name + "-1"
			c <- x
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for _, x := range xs {
			x.name = x.name + "-2"
			c <- x
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()
	fmt.Println("___________________")

	// // for i := 0; i < 20; i++ {
	// // 	fmt.Println(<-c)
	// // }

	for n := range c {
		fmt.Println(n)
	}

}
