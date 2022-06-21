package main

import "fmt"

func main() {
	TwoSum([]int{2, 1, 3}, 4)
}

func TwoSum(numbers []int, target int) [2]int {
	var o [2]int
	for i, v := range numbers {
		l := len(numbers) - 1
		for j := i + 1; j <= l; j += 1 {
			if v+numbers[j] == target {
				o[0] = i
				o[1] = j
				fmt.Println(o)
				return o
			}
		}
	}
	return o
}
