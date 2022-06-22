package main

import (
	"fmt"
	"strings"
)

func main() {
	i := "733049910872815764"
	fmt.Println(Revrot(i, 5))
}

func Revrot(s string, n int) string {
	// deal with edge cases
	if n > len(s) || n <= 0 || len(s) == 0 {
		return ""
	}
	// evaluate all chunks
	sl := []string{}
	for len(s) >= n {
		sl = append(sl, s[:n])
		fmt.Println(sl)
		s = s[n:]
		fmt.Println(s)
	}
	// run chunk transforms
	for i, v := range sl {
		sl[i] = ChunkTransform(v)
	}
	return strings.Join(sl, "")
}

func Rot(s string) string {
	s = s[1:] + string(s[0])
	return s
}

func Rev(s string) string {
	var o string
	for _, v := range s {
		o = string(v) + o
	}
	return o
}

func ChunkTransform(s string) string {
	rt := 0
	for _, v := range s {
		n := int(v)
		rt += (n * n * n)
		fmt.Println(rt)
	}
	if rt%2 == 0 && rt > 0 {
		fmt.Println("Reversing chunk to", Rev(s))
		return Rev(s)
	}
	fmt.Println("Rotating chunk to", Rot(s))
	return Rot(s)
}
