package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("asdksoa", 5)
	want := "asdksoaasdksoaasdksoaasdksoaasdksoa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeatedText := Repeat("Go is awesome! ", 4)
	fmt.Println(repeatedText)

	// Output: Go is awesome! Go is awesome! Go is awesome! Go is awesome!
}
