package main

import (
	"fmt"
)

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  int
	door   int
	color  string
}

func main() {

	v1 := vehicle{
		engine: engine{electric: true},
		make:   "battery body",
		model:  123,
		door:   2,
		color:  "blue",
	}
	v2 := vehicle{
		engine: engine{electric: false},
		make:   "diesel body",
		model:  123,
		door:   2,
		color:  "black",
	}

	fmt.Println(v1)
	fmt.Println(v1.engine.electric)
	fmt.Println(v1.make)
	fmt.Println(v1.model)
	fmt.Println(v1.door)
	fmt.Println(v1.color)
	fmt.Println("_________________________")
	fmt.Println(v2)
	fmt.Println(v2.engine.electric)
	fmt.Println(v2.make)
	fmt.Println(v2.model)
	fmt.Println(v2.door)
	fmt.Println(v2.color)

}
