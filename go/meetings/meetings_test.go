package main

import (
	"strings"
	"testing"
)

func TestStrToSlice(t *testing.T) {
	exp := [][]string{{"WAHL", "ALEXIS"}, {"BELL", "JOHN"}}
	out := strToSlice("Alexis:Wahl;John:Bell")
	for i, v := range exp {
		if strings.Join(v, ",") != strings.Join(out[i], ",") {
			t.Errorf("Expected %v, got %v", v, out[i])
		}
	}
}

func TestSortSlice(t *testing.T) {
	exp := []string{"BELL, JOHN", "WAHL, ALEXIS"}
	out := sortAndCleanSlice([][]string{{"WAHL", "ALEXIS"}, {"BELL", "JOHN"}})
	for i, v := range out {
		if exp[i] != v {
			t.Errorf("Expected %v, got %v", exp[i], v)
		}
	}
}

func TestJoinAndTidy(t *testing.T) {
	exp := "(BELL, JOHN)(WAHL, ALEXIS)"
	out := joinAndTidy([]string{"BELL, JOHN", "WAHL, ALEXIS"})
	if out != exp {
		t.Errorf("Expected %v, but got %v", exp, out)
	}
}
