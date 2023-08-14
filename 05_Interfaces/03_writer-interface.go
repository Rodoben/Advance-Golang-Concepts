package main

import (
	"fmt"
	"io"
	"os"
)

type Person struct {
	name string
	arg  int
	sex  string
}

func (p *Person) WriteOut(w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf("My name is %s and i am %v years old %v", p.name, p.arg, p.sex)))
	return err
}

func main() {
	p := Person{name: "Ronald", arg: 12, sex: "Male"}
	err := p.WriteOut(os.Stdout)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
