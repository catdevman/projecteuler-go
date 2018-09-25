package main

import (
	_ "log"
	"testing"
)

func TestLowestDivisor(t *testing.T) {
	testNumber := 2520
	n := lowestNumberFor(1, 10)

	if n != testNumber {
		t.Fail()
	}
}
