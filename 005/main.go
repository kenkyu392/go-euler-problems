package main

import (
	"fmt"
)

// Problem 5: Smallest multiple
// 2520 is the smallest number that can be divided by each of the numbers
// from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all
// of the numbers from 1 to 20?
func main() {
	const want = 232792560
	got := 1
	for i := 1; i < 21; i++ {
		got = lcm(got, i)
	}
	fmt.Printf("got:%d == want:%d = %t", got, want, got == want)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
