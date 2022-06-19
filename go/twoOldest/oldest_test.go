package main

import "testing"

func TestOldest(t *testing.T) {
	input := []int{10, 12, 16}
	exp := [2]int{12, 16}
	out := oldest(input)
	for i, _ := range out {
		if exp[i] != out[i] {
			t.Errorf("Value mismastch")
		}
	}
}
