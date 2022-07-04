package main

import "fmt"

func main() {
	fmt.Println(ChooseBestSum(163, 3, []int{50, 55, 56, 57, 58}))
	// fmt.Println(SumUp(10, -1, []int{2, 4, 7}))
}

func ChooseBestSum(t, k int, ls []int) int {
	outLs := []int{}
	fmt.Println(outLs)
	outI := -1

	return outI
}

func SumUp(max, cur int, curArr []int) int {
	s := 0
	for _, v := range curArr {
		s += v
	}
	if s > cur && s <= max {
		return s
	}
	return cur
}

func CreateCombinations(arr, data []int, start, end, index, r int) []int {
	if index == r {
		for j := 0; j < r; j++ {
			fmt.Println()
		}
	}

}

// https://www.geeksforgeeks.org/print-all-possible-combinations-of-r-elements-in-a-given-array-of-size-n/
// https://www.codewars.com/kata/55e7280b40e1c4a06d0000aa/train/go
