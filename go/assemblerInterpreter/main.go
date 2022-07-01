package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	i := []string{"mov a -10", "mov b a", "inc a", "dec b", "jnz a -2"}
	fmt.Println(SimpleAssembler(i))
	// fmt.Println(splitter(i[0]))
}

func SimpleAssembler(program []string) map[string]int {
	o := make(map[string]int)
	i := 0
	for i < len(program) {
		so := splitter(program[i], o)
		switch so.cmd {
		case "mov":
			o[so.value] = so.numeric
			i += 1

		case "inc":
			o[so.value] += 1
			i += 1

		case "dec":
			o[so.value] -= 1
			i += 1

		case "jnz":
			if o[so.value] != 0 {
				i += so.numeric
			} else {
				i += 1
			}

		}
		// fmt.Println("o is:", o)
	}
	return o
}

type commands struct {
	cmd     string
	value   string
	numeric int
}

func splitter(s string, o map[string]int) commands {
	sl := strings.Split(s, " ")
	inter := 0
	if len(sl) == 3 {
		if out, _ := regexp.MatchString("[0-9]+", sl[2]); out {
			inter, _ = strconv.Atoi(sl[2])
		} else {
			inter = o[sl[2]]
		}
	}
	c := commands{
		cmd:     sl[0],
		value:   sl[1],
		numeric: inter,
	}
	return c
}
