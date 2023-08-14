package main

import (
	"fmt"
	"strings"
)

func generate(str ...string) <-chan string {
	out := make(chan string)
	go func() {
		for _, s := range str {
			out <- s
		}
		close(out)
	}()
	return out
}

func ispalindrome(c <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for s := range c {
			out <- fmt.Sprintf("Name: %s is palindrome? - %t", s, isPalin(s))
		}
		close(out)
	}()
	return out
}

func isPalin(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	fmt.Println(s)
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func merge(cs ...<-chan string) <-chan string {
	out := make(chan string)
	done := make(chan bool)
	output := func(c <-chan string) {
		for n := range c {
			out <- n
		}
		done <- true
	}
	for _, n := range cs {
		go output(n)
	}

	go func() {
		for range cs {
			<-done
		}
		close(out)
	}()

	return out
}

func fanOut(in <-chan string, n int) []<-chan string {
	var xs []<-chan string
	for i := 0; i < n; i++ {
		xs = append(xs, ispalindrome(in))
	}
	return xs
}

func main() {

	//BAD WAY FOR LAUNCHING GOROUINES

	// c := generate("naman", "rodo", "madam")
	// out := ispalindrome(c)
	// out1 := ispalindrome(c)
	// out2 := ispalindrome(c)
	// out3 := ispalindrome(c)
	// out4 := ispalindrome(c)

	// for n := range merge(out, out1, out2, out3, out4) {
	// 	fmt.Println(n)
	// }

	//------------------------------------------------------------------
	c := generate("naman", "rodo", "madam", "ponoin", "a man nam a", "A man a plan a canal Panama")
	xs := fanOut(c, 100) // here we are fannig out to 10 go routines
	// for _, v := range xs {
	// 	fmt.Println("********", <-v)
	// }

	for r := range merge(xs...) { // Here we are fanning in all the channels to the merge function
		fmt.Println(r)
	}
}
