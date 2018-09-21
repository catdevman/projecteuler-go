package main

import (
	"fmt"
)

//MAX value that is being counted up to
const MAX = 999999999

type result struct {
	divisible bool
	num       int
}

func main() {
	var nums = generateNumbers()
	sum := 0
	for _, i := range nums {
		if divisibleBy3Or5(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}

func generateNumbers() []int {
	var results []int
	for i := 1; i <= MAX; i++ {
		results = append(results, i)
	}
	return results
}

func divisibleBy(divisor int, i int) bool {
	return i%divisor == 0
}

func divisibleBy3Or5(i int) bool {
	return divisibleBy(3, i) || divisibleBy(5, i)
}
