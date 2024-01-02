package main

import "fmt"

var m1 map[string]int

func main() {
	m := map[string]int{
		"ron": 12,
	}

	m["ronald"] = 26
	fmt.Println(m)

	// var m map[string]int{
	// 	"don": 23,
	// 	"derek": 16,
	// } // cant initialize it as such ,  can  initialize the variable globally without trying to insert any element

	m1 = map[string]int{
		"don":   23,
		"derek": 16,
	}
	fmt.Println(m1)

	m2 := make(map[string]int)
	fmt.Println("Zero Value:", m2)
	fmt.Println(len(m2))
	m2["abc"] = 12
	m2["vcd"] = 12
	m2["dfc"] = 12
	m2["wes"] = 12
	fmt.Println(len(m2))
	fmt.Println(m2)

}
