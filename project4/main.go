package main

import (
	"fmt"
	"github.com/hishamkaram/palinagram"
)

const max = 999

func main() {
	hp := highestPalidrone(max, max)
	fmt.Println(hp)
}

func highestPalidrone(first, second int) int {

	n := 0
	h := 0
	for i := first; i > 0; i-- {
		for j := second; j > 0; j-- {
			n = i * j
			if palinagram.IsPalindromeNumber(n) {
				if n > h {
					h = n
				}
			}
		}
	}

	return h
}
