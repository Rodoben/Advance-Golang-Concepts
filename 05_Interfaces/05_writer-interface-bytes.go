package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	b := bytes.NewBufferString("Hello")
	fmt.Println(b)
	b.Write([]byte("Ronald"))
	fmt.Println(b.String()) // stringer interface
	b.Reset()
	b.WriteString("Benjamin")
	fmt.Println(b)
	var r rune = 'A'
	b.WriteRune(r)
	fmt.Println(b)

	f, err := os.Create("file2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	_, err = b.WriteTo(f) // io.writer interface
	if err != nil {
		fmt.Println(err)
	}

}
