package main

import "fmt"

func main() {

	a := [3]int{42, 43, 44}
	fmt.Println(a)

	b := [...]string{"Hello", "Gophers"}
	fmt.Println(b)

	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c)

}
