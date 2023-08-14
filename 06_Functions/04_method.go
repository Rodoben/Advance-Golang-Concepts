package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p *person) Speak() {
	fmt.Printf("I am %s , %v years old", p.first, p.age)
}

func main() {
	p := person{
		first: "Ronald",
		age:   26,
	}

	p.Speak()
}
