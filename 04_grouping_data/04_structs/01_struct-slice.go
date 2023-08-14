package main

import "fmt"

type Person struct {
	first_name string
	last_name  string
	favice     []string
}

func main() {

	p1 := Person{
		first_name: "ronaldf",
		last_name:  "benhjamin",
		favice:     []string{"qwe0", "blaclk"},
	}

	p2 := Person{
		first_name: "derek",
		last_name:  "benhjamin",
		favice:     []string{"qwe0", "blaclk"},
	}
	fmt.Println(p1.first_name, p1.last_name)
	for _, r := range p1.favice {
		fmt.Println(r)
	}
	fmt.Println(p2.first_name, p2.last_name)
	for _, r := range p2.favice {
		fmt.Println(r)
	}
}
