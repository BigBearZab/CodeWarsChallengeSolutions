package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{10, 2}
	fmt.Println(smallest(a))
}

func smallest(na []int) int {
	sort.Slice(na, func(i, j int) bool { return na[i] < na[j] })
	return na[0]
}
