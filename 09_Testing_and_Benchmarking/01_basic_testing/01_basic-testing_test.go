package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum(3, 7)
	want := 10

	if got != want {
		t.Errorf("Error want %v and got %v", want, got)
	}

}
