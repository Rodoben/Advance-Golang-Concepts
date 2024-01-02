package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x) // zero value is [0 0 0 0 0]
	x[3] = 42      // placed at index 3
	fmt.Println(x)
	fmt.Println(len(x))
}
