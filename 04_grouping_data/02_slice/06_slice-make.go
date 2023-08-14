package main

import "fmt"

func main() {
	usa := make([]string, 0, 50)
	fmt.Println(len(usa), cap(usa))

	list := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`,
		`Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
		`Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`,
		`Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`,
		`Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}

	usa = append(usa, list...)

	fmt.Println("_______________________________________")

	// for _, v := range list {
	// 	usa = append(usa, v)
	// }   // Do not uset this , it will affect the memory allocation, 50 cap is needed but this makes it to 112

	fmt.Println(len(usa), cap(usa))

	for i := 0; i < len(usa); i++ {
		fmt.Println(i, usa[i])
	}
	fmt.Println(len(usa), cap(usa))
}
