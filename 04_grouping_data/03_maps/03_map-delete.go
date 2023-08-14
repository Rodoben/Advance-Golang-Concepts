package main

import "fmt"

var m1 map[string]int

func main() {
	m := map[string]int{
		"ron": 12,
	}

	m["ronald"] = 26
	fmt.Println(m)
	delete(m, "ron")
	fmt.Println("Ranging the map m")

	for i, v := range m {
		fmt.Println(i, v)
	}

	// var m map[string]int{
	// 	"don": 23,
	// 	"derek": 16,
	// } // cant initialize it as such ,  do this initializ the variable globally

	m1 = map[string]int{
		"don":   23,
		"derek": 16,
	}
	fmt.Println(m1)
	delete(m1, "don")
	fmt.Println("Ranging the map m1")

	for i, v := range m1 {
		fmt.Println(i, v)
	}

	// m11 := map[int]string{
	// 	0: "zero",
	// 	1: "one",
	// }
	// delete(m11)

	// for i := 0; i < len(m11); i++ {
	// 	fmt.Println(m[i])
	// } // can only range a map but cant loop over it

	m2 := make(map[string]int)
	fmt.Println(len(m2))
	m2["abc"] = 12
	m2["vcd"] = 12
	m2["dfc"] = 12
	m2["wes"] = 12
	fmt.Println(len(m2))
	fmt.Println(m2)
	delete(m2, "abc")
	fmt.Println("Ranging the map m2")

	for i, v := range m2 {
		fmt.Println(i, v)
	}

}
