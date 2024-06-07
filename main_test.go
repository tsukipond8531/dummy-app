package main

import (
	"testing"
)

// TestHello calls main.Hello checking for a valid return value.
func TestHello(t *testing.T) {
	want := "Hello, world"
	msg := Hello()
	if msg != want {
		t.Fatalf("It did not say hello ! How rude !")
	}
}
