package main

import "fmt"

type callbackFunc func(int) int

func iterate(nums []int, cb callbackFunc) {
	for i, num := range nums {
		nums[i] = cb(num)
	}
}

func main() {
	doMath(3, 4, sum) // function call back
	doMath(5, 2, sub) // function call back

	div := func(x, y int) int {
		return x / y
	}

	doMath(5, 2, div) // function call back

	nums := []int{1, 2, 3, 4, 5}
	tripler := func(n int) int {
		return n * 3
	}
	iterate(nums, tripler) // function call back
	fmt.Println(nums)

}

func doMath(x int, y int, f func(x, y int) int) {
	x1 := f(x, y)
	fmt.Println("X1:", x1)
}

func sum(x, y int) int {

	return x + y

}

func sub(x, y int) int {

	return x - y

}
