package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	f := foo(x...)
	fmt.Println(f)
	b := bar(x)
	fmt.Println(b)
}

func foo(v ...int) int { // ...variadic parameter
	var sum int

	for _, v := range v {
		sum += v
	}
	return sum
}
func bar(v []int) int {
	var sum int

	for _, v := range v {
		sum += v
	}
	return sum
}
