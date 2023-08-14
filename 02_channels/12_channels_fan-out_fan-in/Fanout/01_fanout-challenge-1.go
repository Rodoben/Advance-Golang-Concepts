package main

import (
	"fmt"
	"time"
)

var workerId, publisherId int

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
	publisherId++
	thisId := publisherId
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
	workerId++
	thisId := workerId
	for {
		fmt.Printf("%d: waiting for input...\n", thisId)
		input := <-in
		fmt.Printf("%d: input is:%s\n", thisId, input)
	}
}
