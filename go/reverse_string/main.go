package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "word"
	fmt.Println(s)
	fmt.Println(invert(s))
}

func invert(s string) string {
	s_rev := strings.Split(s, "")
	for i, j := 0, len(s_rev)-1; i < j; i, j = i+1, j-1 {
		s_rev[i], s_rev[j] = s_rev[j], s_rev[i]
	}
	return strings.Join(s_rev, "")
}
