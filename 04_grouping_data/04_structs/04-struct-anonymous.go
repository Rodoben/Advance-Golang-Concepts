package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "ronald",
		friends: map[string]int{
			"abc":  1,
			"bvcd": 2,
		},
		favDrinks: []string{"black-current", "ronadl", "rfg"},
	}

	fmt.Println(p1.first)
	for i, v := range p1.friends {
		fmt.Println(i, v)
	}
	for i, v := range p1.favDrinks {
		fmt.Println(i, v)
	}
}
