package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 50; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for n := range ch {
			fmt.Println("From 1:", n)
		}
		wg.Done()
	}()

	go func() {
		for n := range ch {
			fmt.Println("From 2:", n)
		}
		wg.Done()
	}()

	wg.Wait()
}
