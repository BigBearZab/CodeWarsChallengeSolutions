package main

import (
	"testing"
)

func TestPickPeaks(t *testing.T) {
	i1 := []int{0, 1, 2, 4, 3, 1}
	e1 := PosPeaks{[]int{3}, []int{4}}
	o1 := PickPeaks(i1)

	if o1.Pos[0] != e1.Pos[0] && o1.Peaks[0] != e1.Peaks[0] {
		t.Errorf("Expected pos 3, and peak 4, but got %v and %v", e1.Pos[0], e1.Peaks[0])
	}

	i2 := []int{2, 1, 3, 1, 2, 2, 2, 2}
	e2 := PosPeaks{Pos: []int{2}, Peaks: []int{3}}
	o2 := PickPeaks(i2)
	for i, v := range e2.Pos {
		if v != o2.Pos[i] {
			t.Errorf("Expected pos %v, but got pos %v", v, o2.Pos[i])
		}
	}
	if len(o2.Pos) != len(e2.Pos) {
		t.Errorf("Expected length of pos %v, got length %v", len(e2.Pos), len(o2.Pos))
	}

	i3 := []int{2, 1, 3, 1, 2, 2, 2, 2, 1}
	e3 := PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}}
	o3 := PickPeaks(i3)
	for i, v := range e3.Pos {
		if v != o3.Pos[i] {
			t.Errorf("Expected pos %v, but got pos %v", v, o3.Pos[i])
		}
	}
	if len(o3.Pos) != len(e3.Pos) {
		t.Errorf("Expected length of pos %v, got length %v", len(e3.Pos), len(o3.Pos))
	}
}
