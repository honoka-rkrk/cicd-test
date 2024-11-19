package main

import "testing"

func TestEventOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "event" {
		t.Errorf("expected: even,actual: %s", result)
	}
}