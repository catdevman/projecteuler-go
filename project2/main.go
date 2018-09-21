package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	const max = 4000000
	f := fib()
	x := 0
	sum := 0
	for x < max {
		x = f()
		if x < max {
			if x%2 == 0 {
				sum += x
			}
		}
	}
	fmt.Println(sum)
}
