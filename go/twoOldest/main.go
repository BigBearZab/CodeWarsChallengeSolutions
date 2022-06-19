package main

import (
	"fmt"
	"sort"
)

func main() {
	exp := []int{12, 16, 13}
	fmt.Println(oldest(exp))
}

func oldest(ages []int) [2]int {
	sort.Slice(ages, func(i, j int) bool { return ages[i] > ages[j] })
	o := [2]int{ages[1], ages[0]}
	return o
}
