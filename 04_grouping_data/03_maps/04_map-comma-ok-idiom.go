package main

import (
	"fmt"
)

// count the number of vowels and consonants from the given slice

func main() {

	xi := []string{"a", "b", "b", "a", "e", "e", "i", "t"}

	vowels := make(map[string]int)
	consonants := make(map[string]int)

	for _, v := range xi {
		if isvowels(v) {
			if _, ok := vowels[v]; ok { // ok comma idiiom
				vowels[v]++
			} else {
				vowels[v] = 1 // first time register
			}
		} else {
			if _, ok := consonants[v]; ok { // ok comma idiom
				consonants[v]++
			} else {
				consonants[v] = 1 // first time register
			}
		}
	}

	fmt.Println("vowels", vowels)
	fmt.Println("consonants", consonants)
}

func isvowels(s string) bool {
	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
		return true
	}
	return false
}
