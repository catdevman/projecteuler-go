package main

import (
	"fmt"
	"math"
)

func main() {
	r := makeIntRange(1, 100)
	d := squareOfSumForRange(r) - sumOfSquareForRange(r)
	fmt.Println(d)
}

func sumOfSquareForRange(r []int) int {
	sum := 0
	for _, i := range r {
		sum += int(math.Pow(float64(i), float64(2)))
	}
	return sum
}

func squareOfSumForRange(r []int) int {
	sum := 0
	for _, i := range r {
		sum += i
	}
	return int(math.Pow(float64(sum), float64(2)))
}

func makeIntRange(low, high int) []int {
	r := []int{}
	for i := low; i <= high; i++ {
		r = append(r, i)
	}
	return r
}
