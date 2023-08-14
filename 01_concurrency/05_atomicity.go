package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

func main() {

	fmt.Println("using Atomic to solve data race")
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {

	for i := 0; i < 11; i++ {
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter", counter)
	}
	wg.Done()
}
