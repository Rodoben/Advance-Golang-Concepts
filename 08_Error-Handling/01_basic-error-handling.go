package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func init() {
	_, err := os.Open("file.txt")
	if err != nil {
		//fmt.Printf("Error: %v\n\n\n", err)
		//log.Println("Error Unable to marshal json", err)
		//log.Fatalf("Error: %v", err)
		panic(err)
	}
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{`72bhj`, "Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(&p1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		//log.Println("Error Unable to marshal json")
		//log.Fatalf("Error: %v\n", err)
		//panic(err)
	}

	fmt.Println(string(bs))

}
