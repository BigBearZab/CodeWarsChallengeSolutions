package main

import "testing"

func TestSmallest(t *testing.T) {
	a1 := []int{7, 21, 10}
	a2 := []int{1, 14, 6}
	if smallest(a1) != 7 {
		t.Errorf("Expected 7, got %v", smallest(a1))
	}
	if smallest(a2) != 1 {
		t.Errorf("Expected 7, got %v", smallest(a2))
	}
}
