package main

import (
	"fmt"
	"math"
)

// Problem 6: Sum square difference
// The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
// Hence the difference between the sum of the squares of the first ten
// natural numbers and the square of the sum is 3025 − 385 = 2640.
// Find the difference between the sum of the squares of the first one
// hundred natural numbers and the square of the sum.
func main() {
	const N = 100
	const want = 25164150
	/*
		got := 0
		x := 0
		for i := 1; i <= N; i++ {
			got += i * i
			x += i
		}
		got = x*x - got
	*/
	got := int(math.Pow(float64(N*(N+1)/2), 2)) - (N*(N+1)*(2*N+1))/6
	fmt.Printf("got:%d == want:%d = %t", got, want, got == want)
}
