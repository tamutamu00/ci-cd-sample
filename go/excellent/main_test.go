package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "odd" {
		t.Errorf("EvenOrOdd(10) = %s, want odd", result)
	}
}
