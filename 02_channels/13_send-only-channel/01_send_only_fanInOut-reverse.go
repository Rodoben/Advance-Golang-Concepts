package main

import (
	"fmt"
	"time"
)

// reverse the string

func gen(strs ...string) <-chan string {
	out := make(chan string)
	go func() {
		for _, s := range strs {
			out <- s
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan string, n int, out chan<- string) {
	for i := 0; i < n; i++ {
		reverse(in, out)

	}
}

func reverse(c <-chan string, c1 chan<- string) { // send only channel
	go func() {
		for n := range c {
			c1 <- rev(n)
		}
	}()
}

func rev(str string) string {
	s := []byte(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(string(s))
	return string(s)

}

func main() {
	in := gen("ronald", "derek", "madam")
	out := make(chan string)
	fanOut(in, 0, out)

	go func() {
		for v := range out {
			fmt.Println(v)
		}
	}()
	time.Sleep(5 * time.Second)

}
