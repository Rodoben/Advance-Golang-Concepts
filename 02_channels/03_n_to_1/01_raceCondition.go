package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type hotdog struct {
	number int
	name   string
}

func main() {
	wg.Add(3)
	c := make(chan hotdog)

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

	for i, v := range xs {
		v := v
		i := i
		go func(u hotdog, i int) {
			fmt.Println("loop:", i, v)

		}(v, i)
	}
	wg.Done()
	fmt.Println("___________________")
	go func() {
		for _, x := range xs {
			x.name = x.name + "-1"
			c <- x
		}
		wg.Done()
	}()
	fmt.Println("___________________")
	go func() {
		for _, x := range xs {
			x.name = x.name + "-2"
			c <- x
		}
		wg.Done()
	}()
	fmt.Println("___________________")
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
