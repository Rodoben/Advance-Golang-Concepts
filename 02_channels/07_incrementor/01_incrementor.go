package main

import "fmt"

func main() {

	c1 := incrementor("Foo:")
	c2 := incrementor("BAR:")

	c3 := Sumpuller(c1)
	c4 := Sumpuller(c2)
	fmt.Println("Final Counter:", <-c3+<-c4)

	for c5 := range Sumpuller(c2) {
		fmt.Println(c5)
	}

}

//func Sumpuller(c chan int) chan int {
// Here we have example that we can only use bi-directional-channel for typecasting.

func incrementor(s string) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

// add receive only channel here
func Sumpuller(c <-chan int) <-chan int {
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
