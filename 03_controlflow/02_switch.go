package main

import (
	"fmt"
	"math/rand"
)

// Switch case:

func main() {

	x := rand.Intn(296)
	fmt.Println(x)

	switch {
	case x <= 100:
		{
			fmt.Println("Between 0 and 100")
		}

	case x > 100 && x <= 200:
		{

			fmt.Println("Between 100 and 200")
		}
	case x > 200 && x <= 250:
		{
			fmt.Println("Between 200 and 250")

		}
	}

}
