package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where the initialization for my program occurs!")
}

func main() {
	fmt.Println("I am Living inside the main function.")
	x := rand.Intn(298)
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
