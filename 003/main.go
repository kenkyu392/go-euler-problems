package main

import "fmt"

// Problem 3: Largest prime factor
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
func main() {
	const want = 6857
	got := 600851475143
	for i := 3; i*i <= got; i = i + 2 {
		for got%i == 0 {
			got = got / i
		}
	}
	fmt.Printf("got:%d == want:%d = %t", got, want, got == want)
}
