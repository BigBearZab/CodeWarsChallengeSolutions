package main

import "fmt"

func main() {
	fmt.Println(Beeramid(5000, 3))
}

func Beeramid(bonus int, price float64) int {
	maxBeers := int(float64(bonus) / price)
	levelCounter := 0
	beerCounter := 0
	i := 1
	for beerCounter <= maxBeers {
		beerCounter += (i * i)
		if beerCounter > maxBeers {
			break
		}
		levelCounter += 1
		i++
	}
	return levelCounter
}
