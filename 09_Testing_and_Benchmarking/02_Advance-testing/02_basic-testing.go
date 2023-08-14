package main

import "fmt"

type func1 func(n int) []int
type func2 func(n int) int

func main() {

	fmt.Println(FibonacciSeries(10))

	//---------------------------------------------------

	doMath(10, FibonacciSeries)
	//doMath(10, fibonacciSeriesRecursive)
	doMath(25, primeNumbers)

	DoMath2(10, fibonacciSeriesRecursive)
	DoMath2(7, FactorialRecursive)

}

// 0 0 1 2 3 5 8 13 21 34
func FibonacciSeries(n int) []int {
	series := make([]int, n)

	if n >= 1 {
		series[0] = 0
	}

	if n >= 2 {
		series[1] = 1
	}

	for i := 2; i < n; i++ {
		series[i] = series[i-1] + series[i-2]
	}
	return series
}

func fibonacciSeriesRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciSeriesRecursive(n-1) + fibonacciSeriesRecursive(n-2)
}

func primeNumbers(n int) []int {
	primenos := []int{}
	for i := 2; i < n; i++ {
		if isPrime(i) == 1 {
			primenos = append(primenos, i)
		}
	}
	return primenos
}

func isPrime(n int) int {
	if n <= 1 {
		return 0
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return 0
		}
	}
	return 1
}

func FactorialRecursive(n int) int {
	if n <= 0 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

func doMath(n int, f func1) []int {
	f1 := f(n)
	fmt.Println("DoMath:", f1)
	return f1
}

func DoMath2(n int, f func2) []int {

	Series := []int{}

	for i := 0; i < n; i++ {

		Series = append(Series, f(i))
	}

	fmt.Println("Series", Series)
	return Series
}
