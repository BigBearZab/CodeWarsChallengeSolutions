package main

import "testing"

func TestFactorial(t *testing.T) {
	if factorial(0) != 1 {
		t.Errorf("Expected 1, got %v", factorial(0))
	}
	if factorial(2) != 2 {
		t.Errorf("Expected 2, got %v", factorial(2))
	}
	if factorial(10) != 3628800 {
		t.Errorf("Expected 3628800, got %v", factorial(10))
	}
}
