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
	ch := make(chan []Employee)

	go func() {

		ch <- []Employee{
			{name: "Ronald", age: 12, salary: 23.4, contact: map[string]int{
				"primary contact":   3344523,
				"secondary contact": 8736462,
			},
				Address: []*Address{
					{city: "bangalore", state: "Karnataka"},
					{city: "bangalore", state: "Karnataka"},
				},
			},
			{name: "derek", age: 12, salary: 23.4, contact: map[string]int{
				"primary contact":   3344523,
				"secondary contact": 8736462,
			},
				Address: []*Address{
					{city: "bangalore", state: "Karnataka"},
					{city: "bangalore", state: "Karnataka"},
				},
			},
		}
		//close(ch)
	}()

	fmt.Println(<-ch)

	// for v := range ch {
	// 	for _, c := range v {
	// 		for _, v := range c.Address {
	// 			fmt.Println(v.city, v.state)
	// 		}
	// 		for i, v := range c.contact {
	// 			fmt.Println(i, v)
	// 		}
	// 		fmt.Println(c.age)
	// 		fmt.Println(c.contact)
	// 		fmt.Println(c.name)
	// 	}
	// }

	// data := <-ch

	// for i := 0; i < len(data); i++ {
	// 	c := data[i]
	// 	for _, v := range c.Address {
	// 		fmt.Println(v.city, v.state)
	// 	}
	// 	for i, v := range c.contact {
	// 		fmt.Println(i, v)
	// 	}
	// 	fmt.Println(c.age)
	// 	fmt.Println(c.contact)
	// 	fmt.Println(c.name)
	// }
}
