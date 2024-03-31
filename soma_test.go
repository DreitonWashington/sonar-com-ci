package main

import "testing"

func TestSoma(t *testing.T) {
	result := soma(2, 3)
	if result != 5 {
		t.Errorf("Result must be 5")
	}
}
