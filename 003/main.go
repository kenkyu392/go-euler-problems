package main

import "fmt"

// Problem 3: Largest prime factor
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
func main() {
	const want = 6857
	got := 0
	n := 600851475143
	for n%2 == 0 {
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			got = i
			n = n / i
		}
	}
	if n > 2 {
		got = n
	}
	fmt.Printf("got:%d == want:%d = %t", got, want, got == want)
}
