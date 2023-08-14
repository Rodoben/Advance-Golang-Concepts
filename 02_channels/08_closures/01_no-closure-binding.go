package main

import "fmt"
// go vet .
//.\01_no-closure-binding.go:11:29: loop variable v captured by func literal

func main() {
	done := make(chan bool)
	values := []string{"a", "c", "d"}
	for _, v := range values {
		fmt.Println("Instance1:", v)
		go func() {
			fmt.Println("Instance:", v)
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}

}

/*
Some confusion may arise when using closures with concurrency.
One might mistakenly expect to see a, b, c as the output.
What you'll probably see instead is c, c, c. This is because
each iteration of the loop uses the same instance of the variable v,
so each closure shares that single variable. When the closure runs,
it prints the value of v at the time fmt.Println is executed,
but v may have been modified since the goroutine was launched.
To help detect this and other problems before they happen,
run go vet.
SOURCE:
https://golang.org/doc/faq#closures_and_goroutines
*/
