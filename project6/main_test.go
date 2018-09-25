package main

import (
	_ "log"
	"testing"
)

func TestSumOfSquareForRange(t *testing.T) {
	r := makeIntRange(1, 10)

	testNumber := 385
	s := sumOfSquareForRange(r)

	if s != testNumber {
		t.Fail()
	}
}

func TestSquareOfSumForRange(t *testing.T) {
	r := makeIntRange(1, 10)

	testNumber := 3025
	s := squareOfSumForRange(r)

	if s != testNumber {
		t.Fail()
	}
}
