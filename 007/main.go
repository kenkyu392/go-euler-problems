package main

import (
	"fmt"
	"math"
)

// Problem 7: 10001st prime
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
// we can see that the 6th prime is 13.
// What is the 10001st prime number?
func main() {
	const N = 10001
	const want = 104743
	got := 0
	c := 0
	for c < N {
		got++
		if isPrime(got) {
			c++
		}
	}
	fmt.Printf("got:%d == want:%d = %t", got, want, got == want)
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	} else if n < 2 {
		return false
	} else if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
