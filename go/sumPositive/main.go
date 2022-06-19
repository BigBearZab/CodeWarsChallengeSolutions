package main

import "fmt"

func sumPositive(numbers []int) int {
	s := 0
	for _, v := range numbers {
		if v > 0 {
			s += v
		}
	}
	return s
}

func main() {
	sp := sumPositive([]int{1, 0, 3, -2})
	fmt.Println(sp)
}
