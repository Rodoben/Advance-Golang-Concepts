package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5, 67))
}

func Sum(n ...int) int {
	var s int
	for _, v := range n {
		s += v
	}
	return s
}
