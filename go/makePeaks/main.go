package main

import "fmt"

func main() {
	fmt.Println(PickPeaks([]int{6, -1, -2, 5, 13, 6, 9, -2, 3, 3, 5, -5}))
}

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(a []int) PosPeaks {
	var p PosPeaks
	l := len(a)
	p.Pos = []int{}
	p.Peaks = []int{}
	for i := 1; i < l-1; i++ {
		if a[i-1] < a[i] && a[i+1] <= a[i] {
			// iterate over subsequent to check if true peak
			for j := i; j < l; j++ {
				if a[j] < a[i] {
					p.Pos = append(p.Pos, i)
					p.Peaks = append(p.Peaks, a[i])
					break
				} else if a[j] > a[i] {
					break
				}
			}

		}
	}
	return p
}
