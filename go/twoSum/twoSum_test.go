package main

import "testing"

func TestTwoSum(t *testing.T) {
	e := [2]int{0, 2}
	i1 := []int{1, 2, 3}
	i2 := 4
	o := TwoSum(i1, i2)
	for i, v := range e {
		if o[i] != v {
			t.Errorf("Expected %v at position %v, but got %v", v, i, o[i])
		}
	}
}
