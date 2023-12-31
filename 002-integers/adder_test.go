package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(5, 5)
	fmt.Println(sum)
	// Output: 10
}
