package main

import "fmt"

//FAN IN
//
//-----c1----|
//------c2---|-----c(n(1+2+3))-------------
//-------c3--|
//
/*
FAN OUT
Multiple functions reading from the same channel until that channel is closed
FAN IN
A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.
*/

func main() {
	c := fanIn(boring("ron"), boring("don"), boring("dek"))
	// for n := range c {
	// 	fmt.Println(n)
	// }
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; ; i++ {
			out <- fmt.Sprintf("%s:%v", msg, i)
		}
		close(out)
	}()

	return out
}

func fanIn(input1, input2, input3 <-chan string) <-chan string {

	out := make(chan string)

	go func() {
		for {
			out <- <-input1
		}
	}()

	go func() {
		for {
			out <- <-input2
		}
	}()

	go func() {
		for {
			out <- <-input3
		}
	}()
	return out

}
