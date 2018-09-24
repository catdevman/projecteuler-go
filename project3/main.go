package main

import "fmt"

// Get all prime factors of a given number n
func primeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func highestPrimeFactor(pfs []int) int {
	hpf := 0
	for _, i := range pfs {
		if i > hpf {
			hpf = i
		}
	}
	return hpf
}

func main() {
	pfs := primeFactors(600851475143)
	hpf := highestPrimeFactor(pfs)

	fmt.Println(hpf)

}
