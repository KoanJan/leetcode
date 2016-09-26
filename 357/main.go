package main

import (
	"fmt"
)

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	// calc
	var (
		c   int = 9
		res int = 9
		i   int = n - 1
	)
	for i > 0 {
		res *= c
		c -= 1
		i -= 1
	}
	return res + countNumbersWithUniqueDigits(n-1)
}

func main() {
	// test
	n := 3
	fmt.Printf("countNumbersWithUniqueDigits(%d) = %d\n", n, countNumbersWithUniqueDigits(n))
}
