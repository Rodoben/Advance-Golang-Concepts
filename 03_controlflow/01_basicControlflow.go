package main

import (
	"fmt"
	"math/rand"
)

// Sequential and Conditonal control flow

func main() {
	x := rand.Intn(240)

	fmt.Println(x)

	if x < 100 {
		fmt.Println("between 0 and 100")
	}

	if x > 100 && x < 200 {
		fmt.Println("between 101 and 200")
	}

	if x > 201 && x < 250 {
		fmt.Println("between 201 and 250")
	}

	y := rand.Intn(1)

	fmt.Println(y)

}
