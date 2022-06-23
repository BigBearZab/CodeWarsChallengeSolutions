package main

import "testing"

func TestSliceMutator(t *testing.T) {
	sl_exp := []int{1, 2, 3}
	sl_rec := []int{1, 2, 3, 4}
	Slicetest(sl_exp, sl_rec, t)
}
