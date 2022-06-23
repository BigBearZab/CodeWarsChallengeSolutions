package main

// import "slicetest"

func main() {
	// run slice test
	s := []int{1, 2, 3}
	s = MutateSlice(s)
	PrintSlice(s)
}

func MutateSlice(sl []int) []int {
	sl[0] = sl[0] + 10
	return sl
}
