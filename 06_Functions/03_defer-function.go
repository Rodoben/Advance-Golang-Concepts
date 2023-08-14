package main

import "fmt"

func main() {

	func1()
	defer func2()
	func3()
	defer func4()
	func5()
}

func func1() {
	fmt.Println("I am Function 1")

}
func func2() {
	fmt.Println("I am Function 2")
	defer func() {
		fmt.Println("I am Function anon func under 2")
	}()

}
func func3() {
	fmt.Println("I am Function 3")

}
func func4() {
	fmt.Println("I am Function 4")

}
func func5() {
	fmt.Println("I am Function 5")

	defer func() {
		fmt.Println("I am Function anon func under 5")
	}()

}
