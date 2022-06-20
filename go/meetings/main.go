package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fmt.Println(meeting("Alexis:Wahl;John:Bell;Bravo:Alpha"))
	// fmt.Println(a)
}

func strToSlice(s string) [][]string {
	sl := strings.Split(s, ";")
	var sl_out [][]string
	for _, v := range sl {
		out := strings.Split(strings.ToUpper(v), ":")
		first, last := out[0], out[1]
		sl_out = append(sl_out, []string{last, first})
	}
	return sl_out
}

func sortAndCleanSlice(ss [][]string) []string {
	var so []string
	for _, v := range ss {
		s := strings.Join(v, ", ")
		so = append(so, s)
	}
	sort.Strings(so)
	return so
}

func joinAndTidy(ss []string) string {
	for i, v := range ss {
		ss[i] = "(" + v + ")"
	}
	s := strings.Join(ss, "")
	return s
}

func meeting(s string) string {
	sls := strToSlice(s)
	so := sortAndCleanSlice(sls)
	sfinal := joinAndTidy(so)
	return sfinal
}
