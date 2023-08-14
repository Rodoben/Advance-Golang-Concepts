package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	name string
	age  int
}

func (p *Person) WriteOut(w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf("I am a Person Named %s and age %v:", p.name, p.age)))
	return err
}

func main() {

	f, err := os.Create("file3.txt")
	if err != nil {
		log.Fatalf("", err)
	}
	defer f.Close()

	p := Person{name: "Ronald", age: 12}
	err = p.WriteOut(f) // writing to a file

	var b bytes.Buffer
	fmt.Println("Writeout")
	p.WriteOut(&b) // writing to a buffer
	fmt.Println("close")
	fmt.Println(b.String())
	b.Reset()
	fmt.Println(b.String())
	p.WriteOut(&b)
	fmt.Println(b.String())
}
