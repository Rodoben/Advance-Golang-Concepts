package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var workerId, publisherId int64

func main() {
	input := make(chan string)

	go workerProcess(input)
	// go workerProcess(input)
	// go publisher(input)
	go publisher(input)
	time.Sleep(1 * time.Millisecond)
	fmt.Println("________________________")

	// for n := range input {
	// 	fmt.Println("range", n)
	// }
}

func publisher(out chan string) {

	atomic.AddInt64(&publisherId, 1)
	thisId := atomic.LoadInt64(&publisherId)

	// publisherId++
	// thisId := publisherId
	dataId := 0
	for {
		dataId++
		fmt.Printf("Publisher %d is pushing data \n", thisId)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisId, dataId)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	fmt.Println("IN:", in)
	atomic.AddInt64(&workerId, 1)
	thisId := atomic.LoadInt64(&workerId)
	for {
		fmt.Printf("%d: waiting for input...\n", thisId)
		input := <-in
		fmt.Printf("%d: input is:%s\n", thisId, input)
	}
}
