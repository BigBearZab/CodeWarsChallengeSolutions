package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := "1234"
	fmt.Println(Revrot(i, 3))
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
	for _, v := range sl {
		temp_n, _ := strconv.ParseInt(v, 0, 0)
		fmt.Println(temp_n)

	}
	return ""
}

// func Rot (s string) string {

// }
