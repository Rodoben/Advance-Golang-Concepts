package main

import "fmt"

func main() {
	f := foo()()
	fmt.Println(f)
}

func foo() func() int {
	x := 3 // function closure
	return func() int {
		x = 9
		x++
		return x
	}
}
