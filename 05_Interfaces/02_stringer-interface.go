package main

import (
	"fmt"
	"log"
)

type car struct {
	name  string
	wheel int
	seats int
	doors int
}

type bike struct {
	name   string
	wheels int
	seats  int
}

func (c *car) String() string {
	return fmt.Sprintf("CarName:%s No_of_wheels: %v No of seats:%v No of doors:%v", c.name, c.wheel, c.seats, c.doors)
}
func (b *bike) String() string {
	return fmt.Sprintf("BikeName:%s No_of_wheels: %v No of seats:%v", b.name, b.wheels, b.seats)
}

func Result(r fmt.Stringer) {
	log.Println("LOG FROM LINE:", r.String())
}

func main() {

	c := car{name: "BMW", wheel: 4, seats: 4, doors: 4}
	b := bike{name: "Hero Honda", wheels: 4, seats: 2}

	fmt.Println(c)
	fmt.Println(b)

	Result(&c)
	Result(&b)

}
