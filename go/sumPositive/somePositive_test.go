package main

import "testing"

func TestSumPositive(t *testing.T) {
	if sumPositive([]int{1, 0, 3, -2}) != 4 {
		t.Errorf("Incorrect sum returned")
	}
}
