package main

import "fmt"

func main() {
	done := make(chan bool)
	values := []string{"a", "c", "d"}
	for _, v := range values {
		//fmt.Println("Instance1:", v)
		v := v
		go func() {
			fmt.Println("Instance:", v)
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}

}
