package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hamming(10))
}

func hamming(n int) uint {
	i, j, k := 0, 0, 0
	x2, x3, x5 := 2.0, 3.0, 5.0
	sl := []float64{1}
	// r := math.Pow(float64(2), i) * math.Pow(float64(3), j) * math.Pow(float64(5), k)
	for c := 1; c < n; c++ {
		sl = append(sl, math.Min(x2, math.Min(x3, x5)))
		// fmt.Println(sl)
		if sl[c] == x2 {
			i++
			x2 = 2 * sl[i]
		}
		if sl[c] == x3 {
			j++
			x3 = 3 * sl[j]
		}
		if sl[c] == x5 {
			k++
			x5 = 5 * sl[k]
		}
	}
	fmt.Println(sl)
	return uint(sl[n-1])
}

// 0 0 0 = 1
// 1 0 0 = 2
// 0 1 0 = 3
// 2 0 0 = 4
// 0 0 1 = 5
// 1 1 0 = 6
// 3 0 0 = 8
// 0 2 0 = 9
// 1 0 1 = 10
// 2 1 0 = 12
// 0 1 1 = 15
// 4 0 0 = 16
// 1 2 0 = 18
// 2 0 1 = 20
