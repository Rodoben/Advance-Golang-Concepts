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
		{number: 1, name: "Gorotine"},
		{number: 2, name: "Gorotine"},
		{number: 3, name: "Gorotine"},
		{number: 4, name: "Gorotine"},
		{number: 5, name: "Gorotine"},
		{number: 6, name: "Gorotine"},
		{number: 7, name: "Gorotine"},
		{number: 8, name: "Gorotine"},
		{number: 9, name: "Gorotine"},
		{number: 10, name: "Gorotine"},
		{number: 11, name: "Gorotine"},
	}

	for i, v := range xs {
		v := v
		i := i
		go func(u hotdog, i int) {
			fmt.Println("loop", i, v)

		}(v, i)

	}
	wg.Done()
	go func() {
		for _, x := range xs {
			x.name = x.name + "-1"
			c <- x
		}
		wg.Done()

	}()

	go func() {
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

	// for i := 0; i < 20; i++ {
	// 	fmt.Println(<-c)
	// }

	for n := range c {
		fmt.Println(n)
	}

}
