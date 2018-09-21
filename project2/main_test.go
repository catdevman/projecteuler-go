package main

import (
	_ "log"
	"testing"
)

func TestFib(t *testing.T) {
	testSeq := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	var fibSeq []int
	f := fib()
	for _ = range [10]int{} {
		fibSeq = append(fibSeq, f())
	}

	if len(testSeq) != len(fibSeq) {
		t.Fail()
	}
	for i, v := range testSeq {
		if v != fibSeq[i] {
			t.Fail()
		}
	}

}
