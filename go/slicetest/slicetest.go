package main

import (
	"fmt"
	"testing"
)

func PrintSlice(i []int) {
	fmt.Println(i)
}

// create a common way to test slices
func Slicetest(sl_exp, sl_rec []int, t *testing.T) *testing.T {
	// start by validating length
	if len(sl_exp) != len(sl_rec) {
		t.Errorf("Expected slice of length %v, got slice of length %v instead", len(sl_exp), len(sl_rec))
	}
	// now iterate over partition to match vales
	for i, v := range sl_exp {
		if v != sl_rec[i] {
			t.Errorf("Expected %v at position %v, but got %v instead", v, i, sl_rec[i])
		}
	}
	return t
}
