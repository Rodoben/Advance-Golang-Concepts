package main

import (
	"fmt"

	"testing"
)

func TestFibonacciSeries(t *testing.T) {

	n := 10
	got := FibonacciSeries(n)
	want := func(n int) []int {
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
		//series[5] = 90
		return series
	}(10)

	if len(got) != len(want) {
		t.Errorf("Error: want %v got %v", want, got)
		fmt.Println(got, "|", want)
		// fmt.Println(len(got), len(want))
	}
	for i := 0; i < n; i++ {
		if got[i] != want[i] {
			t.Errorf("Error in the values want %v got %v", want, got)
		}
	}
}

func TestFibonacciSeriesRecursive(t *testing.T) {
	n := 5
	got := []int{}
	for i := 0; i < n; i++ {
		got = append(got, fibonacciSeriesRecursive(i))
	}
	//got[3] = 12  to show error
	want := FibonacciSeries(n)

	if len(got) != len(want) {
		t.Errorf("Error in recursion: got %v want %v", got, want)
	}
	for i := 0; i < n; i++ {
		if got[i] != want[i] {
			t.Errorf("Error in the values want %v got %v", want, got)
		}
	}
}

func TestPrimeNumbers(t *testing.T) {
	n := 101
	got := primeNumbers(n)
	want := []int{}

	func() {
		for i := 0; i < n; i++ {
			if isPrime(i) == 1 {
				want = append(want, i)
			}
		}
	}()

	fmt.Println(len(got), len(want))
	minLen := len(got)
	if len(want) < minLen {
		minLen = len(want)
	}

	if len(got) != len(want) {
		t.Errorf("Error in Prime: got %v want %v", got, want)
	}

	for i := 0; i < minLen; i++ {
		if got[i] != want[i] {
			t.Errorf("Error in the Primevalues want %v got %v", want, got)
		}
	}

}

func TestIsPrime(t *testing.T) {
	n := 7
	got := isPrime(n)
	want := 1
	if got != want {
		t.Errorf("Error in Prime: got %v want %v", got, want)
	}

}

func TestFactorialRecursive(t *testing.T) {
	//series := []int{}
	n := 4

	got := FactorialRecursive(n)
	want := 24

	if got != want {
		t.Errorf("Error in FactorialRecursive: got %v want %v", got, want)
	}
}

type work struct {
	got  []int
	want []int
}

func TestDoMath_TableTest(t *testing.T) {
	n := 4

	w := []work{
		{
			got:  doMath(n, primeNumbers),
			want: primeNumbers(n),
		},
		{
			got:  doMath(n, FibonacciSeries),
			want: FibonacciSeries(n),
		},
	}

	for _, v := range w {
		if len(v.got) != len(v.want) {
			t.Errorf("Error in recursion: got %v want %v", v.got, v.want)
		}
		for i := 0; i < len(v.got); i++ {
			if v.got[i] != v.want[i] {
				t.Errorf("Error in the values want %v got %v", v.want, v.got)
			}
		}
	}

}

func TestDoMath2_TableTest(t *testing.T) {
	n := 4

	want1 := func(n int) []int {
		w1 := []int{}
		for i := 0; i < n; i++ {
			w1 = append(w1, FactorialRecursive(i))
		}
		return w1
	}(n)

	want2 := func(n int) []int {
		w1 := []int{}
		for i := 0; i < n; i++ {
			w1 = append(w1, fibonacciSeriesRecursive(i))
		}
		return w1
	}(n)

	w := []work{

		{
			got:  DoMath2(n, FactorialRecursive),
			want: want1,
		},

		{
			got:  DoMath2(n, fibonacciSeriesRecursive),
			want: want2,
		},
	}

	for _, v := range w {
		if len(v.got) != len(v.want) {
			t.Errorf("Error in recursion: got %v want %v", v.got, v.want)
		}
		for i := 0; i < len(v.got); i++ {
			if v.got[i] != v.want[i] {
				t.Errorf("Error in the values want %v got %v", v.want, v.got)
			}
		}
	}
}
