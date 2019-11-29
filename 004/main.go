package main

import (
	"fmt"
	"strconv"
)

// Problem 4: Largest palindrome product
// A palindromic number reads the same both ways. The largest palindrome
// made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
func main() {
	const want = 906609
	got := 0
	for x := 1000; x > 100; x-- {
		for y := 1000; y > 100; y-- {
			z := x * y
			if isPalindrome(strconv.Itoa(z)) {
				got = z
			}
			if z < got {
				break
			}
		}
	}
	fmt.Printf("got:%d == want:%d = %t", got, want, got == want)
}

func isPalindrome(s string) bool {
	return len(s)%2 == 0 && s[:len(s)/2] == reverse(s[len(s)/2:])
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
