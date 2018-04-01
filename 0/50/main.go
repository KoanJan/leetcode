package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	var (
		nIsMinus bool    = n < 0
		res      float64 = 1.0
	)
	if nIsMinus {
		n = -n
	}
	for n >= 2 {
		if n%2 == 1 {
			res *= x
		}
		x *= x
		n /= 2
	}
	if n == 1 {
		res *= x
	}
	if nIsMinus {
		return 1.0 / res
	}
	return res
}

func main() {
	// test
	var (
		x float64 = 8.88023
		n int     = 8
	)
	fmt.Printf("myPow(%8f, %d) = %16f\n", x, n, myPow(x, n))
}
