package main

import "fmt"

type Address struct {
	city  string
	state string
}

type Employee struct {
	name    string
	age     int
	salary  float64
	contact map[string]int
	Address []*Address
}

func main() {
	ch := make(chan Employee)

	go func() {
		ch <- Employee{name: "Ronald", age: 12, salary: 23.4, contact: map[string]int{
			"primary contact":   3344523,
			"secondary contact": 8736462,
		},
			Address: []*Address{
				{city: "bangalore", state: "Karnataka"},
				{city: "bangalore", state: "Karnataka"},
			},
		}
	}()

	fmt.Println(<-ch)

}
