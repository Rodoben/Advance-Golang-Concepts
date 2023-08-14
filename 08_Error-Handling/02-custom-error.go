package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{`72bhj`, "Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJson(&p1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		//log.Println("Error Unable to marshal json")
		//log.Fatalf("Error: %v\n", err)
		//panic(err)
	}

	fmt.Println(string(bs))

}

func toJson(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(&a)
	if err != nil {
		//return []byte{}, fmt.Errorf("There was error in toJson %v", err)
		return []byte{}, errors.New("There was error in toJson")
	}
	return bs, nil
}
