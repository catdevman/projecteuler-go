package main

import (
	"fmt"
)

func main() {
	num := lowestNumberFor(1, 20)
	fmt.Println(num)
}

func lowestNumberFor(low, high int) int {
	r := []int{}
	for i := low; i <= high; i++ {
		r = append(r, i)
	}
	num := 1
	for {
		if checkDivisbleEvenlyForRange(num, r) {
			break
		}
		num++
	}
	return num
}

func checkDivisbleEvenlyForRange(n int, r []int) bool {
	for _, i := range r {
		if n%i != 0 {
			return false
		}
	}
	return true
}
