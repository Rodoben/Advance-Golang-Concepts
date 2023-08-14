package main

import (
	"fmt"
	"math/rand"
)

func main() {

	r1 := rand.Intn(10)
	r2 := rand.Intn(10)

	fmt.Println("R1:", r1)
	fmt.Println("R2:", r2)

	if r1 < 4 && r2 < 4 {
		fmt.Println("Both are less than 4")
	}

	if r1 > 6 && r2 > 6 {
		fmt.Println("Both are greater than 6")
	}

	if r1 >= 4 && r2 <= 6 {
		fmt.Println("ranges 4 to 6")
	}

	if r2 != 5 {
		fmt.Println("r2 is not 5")
	}

	if r1 > 10 && r2 > 10 {
		fmt.Println("none of the previous cases are met!")
	}

	switch {
	case r1 < 4 && r2 < 4:
		fmt.Println("Both are less than 4")
	case r1 > 6 && r2 > 6:
		fmt.Println("Both are greater than 6")
	case r1 >= 4 && r2 <= 6:
		fmt.Println("ranges 4 to 6")
	case r2 != 5:
		fmt.Println("r2 is not 5")
	case r1 > 10 && r2 > 10:
		fmt.Println("none of the previous cases are met!")
	default:
		fmt.Println("I am out")
	}

}
