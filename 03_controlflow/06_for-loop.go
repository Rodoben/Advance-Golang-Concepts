package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {
		r1 := rand.Intn(5)

		if r1 < 4 {
			fmt.Println("Both are less than 4")
		}

		if r1 > 6 {
			fmt.Println("Both are greater than 6")
		}

		if r1 >= 4 {
			fmt.Println("ranges 4 to 6")
		}

		if r1 != 5 {
			fmt.Println("r2 is not 5")
		}
	}

}
