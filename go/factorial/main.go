package main

func main() {

}

func factorial(n int) int {
	s := 1
	for n > 1 {
		s *= n
		n -= 1
	}
	return s
}
