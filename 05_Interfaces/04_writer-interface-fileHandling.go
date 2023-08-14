package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.Create("file1.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()
	s := []byte("hello Gophers!")
	_, err = f.Write(s)
	if err != nil {
		log.Println(err)
	}

}
