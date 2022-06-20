package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	s1 := "12.257.56.1"
	fmt.Println(ip_validation(s1))
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re.FindAllString(s1, 4))
	conv_str, _ := strconv.ParseInt("25", 0, 0)
	fmt.Println(conv_str < 256)
}

func ip_validation(ip string) bool {
	re := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
	if re.Match([]byte(ip)) {
		re2 := regexp.MustCompile(`\d+`)
		octets := re2.FindAllString(ip, 4)
		for _, v := range octets {
			re3 := regexp.MustCompile(`0$|^[1-9]\d*`)
			if !re3.Match([]byte(v)) {
				return false
			}
			conv_str, _ := strconv.ParseInt(v, 0, 0)
			if conv_str > 255 {
				return false
			}
		}
		return true
	}
	return false
}
