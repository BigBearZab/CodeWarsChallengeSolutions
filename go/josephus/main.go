package main

import "fmt"

func main() {
	i := []interface{}{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Josephus(i, 3))
}

func Josephus(items []interface{}, k int) []interface{} {
	if len(items) == 0 {
		return items
	}
	o := []interface{}{}
	c := k
	for len(items) > 1 {
		i := circularLoop(items, c)
		fmt.Println("i, c are at", i, c)
		o = append(o, items[i])
		fmt.Println("o is at", o)
		items = RemoveIndex(items, i)
		fmt.Println("items is", items)
		c = i + k
	}

	return append(o, items[0])
}

func circularLoop(s []interface{}, n int) int {
	// take a slice and return nth element where n is the total number passed in, even if n > len slice
	// var r int
	r := (n % len(s)) - 1
	if r == -1 {
		return len(s) - 1
	}
	return r
}

func RemoveIndex(s []interface{}, index int) []interface{} {
	ret := make([]interface{}, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
