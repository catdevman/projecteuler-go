package main

import (
	_ "log"
	"testing"
)

func TestMerge(t *testing.T) {
	in1 := make(chan int)
	in2 := make(chan int)
	go func() {
		in1 <- 1
		in1 <- 2
		in2 <- 3
		in2 <- 4
		close(in1)
		close(in2)
	}()
	count := 0
	for _ = range merge(in1, in2) {
		count++
	}

	if count != 4 {
		t.Fail()
	}
}

func TestDivisibleBy(t *testing.T) {
	first := divisibleBy(3, 3)
	if first != true {
		t.Fail()
	}

	second := divisibleBy(5, 5)
	if second != true {
		t.Fail()
	}
}
