package main

import "fmt"

func main() {
	xi := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "I'm 008."},
	}

	for _, v := range xi {
		fmt.Println(v)
	}
	fmt.Println("__________________________")

	for _, v := range xi {
		for _, v1 := range v {
			fmt.Println(v1)
		}

	}
}
