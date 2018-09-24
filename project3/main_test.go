package main

import (
	_ "log"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	expectedPfs := []int{5, 7, 13, 29}
	pfs := primeFactors(13195)

	if len(expectedPfs) != len(pfs) {
		t.Fail()
	}
	for i, v := range expectedPfs {
		if v != pfs[i] {
			t.Fail()
		}
	}
}

func TestHighestPrimeFactors(t *testing.T) {
	expectedHighest := 29
	pfs := primeFactors(13195)
	hpf := highestPrimeFactor(pfs)

	if hpf != expectedHighest {
		t.Fail()
	}

	expectedHighest = 6857
	pfs = primeFactors(600851475143)
	hpf = highestPrimeFactor(pfs)

	if hpf != expectedHighest {
		t.Fail()
	}
}
