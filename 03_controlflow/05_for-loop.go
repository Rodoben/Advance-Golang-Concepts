package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	x := rand.Intn(296)
	fmt.Println(x)

	switch {
	case x <= 100:
		{
			fmt.Println("Between 0 and 100")
			for i := 0; i < 100; i++ {
				fmt.Println(i)
			}
		}

	case x > 100 && x <= 200:
		{
			fmt.Println("Between 101 and 200")
			for i := 101; i < 201; i++ {
				fmt.Println(i)
			}
		}
	case x > 200 && x <= 250:
		{
			fmt.Println("Between 200 and 250")
			for i := 201; i < 250; i++ {
				fmt.Println(i)
			}

		}
	}

}
