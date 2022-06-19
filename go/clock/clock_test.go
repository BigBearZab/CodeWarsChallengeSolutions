package main

import "testing"

func TestClock(t *testing.T) {
	h, m, s := 0, 0, 0
	if clock(h, m, s) != 0 {
		t.Errorf("Expected %v, got %v", 0, clock(h, m, s))
	}

	if clock(0, 0, 1) != 1000 {
		t.Errorf("Expected %v, got %v", 1000, clock(h, m, s))
	}

	if clock(0, 1, 1) != 61000 {
		t.Errorf("Expected %v, got %v", 61000, clock(h, m, s))
	}
}
