package main

import "fmt"

type Dog struct {
	first string
}

func (d Dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walk.")
}
func (d *Dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {

	d1 := Dog{"Henry"}
	youngRun(&d1)

	d2 := Dog{"Padet"}
	youngRun(&d2)

	d2 = d2.changeName("Rover")
	youngRun(&d2)
}

func (d Dog) changeName(s string) Dog {

	d.first = s
	return d

}
