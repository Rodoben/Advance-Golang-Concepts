package main

import (
	"fmt"

	pup "github.com/Rodoben/puppy"
)

func main() {

	p1 := pup.Bark()
	fmt.Println(p1)
	fmt.Println(pup.BigBarks())
	fmt.Println(pup.SmallBarks())
	pup.From11()
	pup.From12()
	pup.From13()
}
