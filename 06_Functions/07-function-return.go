package main

import "fmt"

func main() {

	f := foo("ronald", "benjamin")
	fmt.Println(f("ronald", "benjamin", "abc"))

	b := bar()
	fmt.Println(b(32))

}

func foo(x, y string) func(x, y, z string) string {
	fmt.Println("I have execued this line too")
	return func(x, y, z string) string {
		return fmt.Sprintf("I am %s%s%s", x, y, z)
	}
	// i dont have access to go here
}

func bar() func(x int) int {
	fmt.Println("i have bar func priveledge")
	return func(x int) int {
		fmt.Println("inside func 1", x)
		return x
	}

	// I cannot do anything here also does not matter if i do anything here
	x := 0
	fmt.Println("mid", x)

	return func(x int) int {
		fmt.Println("inside func 2", x)
		return x
	}

}
