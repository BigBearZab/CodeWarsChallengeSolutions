package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(ValidISBN10("048665088x"))
	// for i :=
}

func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		fmt.Println("len:", len(isbn))
		return false
	}
	sum := 0
	for i := 1; i <= 9; i++ {
		l := string(isbn[i-1])
		o, _ := regexp.MatchString("[0-9]", l)
		if !o {
			return false
		}
		t, _ := strconv.Atoi(l)
		sum += t * i

		fmt.Println(i, sum, t)
	}
	if string(isbn[9]) == "X" || string(isbn[9]) == "x" {
		sum += 100
		fmt.Println(sum)
	} else {
		t, _ := strconv.Atoi((string(isbn[9])))
		sum += t * 10
	}
	fmt.Println(sum % 11)
	return sum%11 == 0
}
