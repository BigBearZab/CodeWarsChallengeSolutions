package main

import "fmt"

func main() {
	fmt.Println(clock(0, 0, 0))
}

func clock(h, m, s int) int {
	seconds := h*3600 + m*60 + s
	return seconds * 1000
}
