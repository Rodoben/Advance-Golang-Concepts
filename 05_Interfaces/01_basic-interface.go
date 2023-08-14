package main

import (
	"fmt"
	"math"
)

// if two types implements the method of an interface then the interface type is also of same TYPE

type circle struct {
	radius float64
}

type square struct {
	length, breadth float64
}

type shape interface {
	area(s string) float64
}

func (c *circle) area(s string) float64 {
	fmt.Printf("I am running :===%s===>: operation\n", s)
	return math.Pi * c.radius * c.radius
}

func (s *square) area(str string) float64 {

	fmt.Printf("I am running :===%s===>: operation\n", str)
	return s.length * s.breadth

}

func (s *square) areaOne(str string) {
	fmt.Printf("I am of type %T and I am %s\n", s, str)
}

func Output(s shape, str string) {
	fmt.Println("______________________________")
	fmt.Printf("I am of same TYPE :===%T===>:\n", s)
	fmt.Println(s.area(str))

}

func main() {
	c := circle{radius: 12.3}
	c.area("Circle")
	s := square{length: 12.3, breadth: 12.3}
	s.area("Square")
	//s.area("Square").areaOne()
	//c.areaOne()
	s.areaOne("areaone attach")

	Output(&c, "Circle")
	Output(&s, "Square")

}
