package main

import "fmt"

type Contact struct {
	phno  int
	email string
	addr  string
}

func (c *Contact) sendNotifcation(p Person) {
	fmt.Printf("Notification pushed on mobile to %v on phno %v:\n", p.name, c.phno)
	fmt.Printf("Notification pushed on email to %v on phno %v:\n", p.name, c.email)
	fmt.Printf(" card Delivered  on Addr to %v on phno %v:\n", p.name, c.addr)
	fmt.Printf(" user %v prime %v:\n", p.name, p.vip)
}

type Person struct {
	name    string
	cards   []string
	vip     bool
	contact Contact
}

func main() {

	c1 := Contact{
		phno:  9986,
		email: "dodo.ben.rb@gmail.com",
		addr:  "belabagan",
	}

	p1 := Person{

		name:  "ronald benjamin",
		cards: []string{"visa", "mastero"},
		vip:   true,
		contact: Contact{
			phno:  9986,
			email: "dodo.ben.rb@gmail.com",
			addr:  "belabagan",
		},
	}

	m1 := make(map[string]Person)
	m1 = map[string]Person{
		"ronald": p1,
	}

	c1.sendNotifcation(p1)

	for i, v := range m1 {
		fmt.Println(i, v)
		for k, v1 := range v.cards {
			fmt.Println(k, v1)
		}

		fmt.Println(v.contact.phno)

	}

}
