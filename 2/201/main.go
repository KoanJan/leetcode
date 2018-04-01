package main

import (
	"fmt"
)

func rangeBitwiseAnd(m int, n int) int {
	// fmt.Printf("m: %b n: %b\n", m, n)
	var (
		d1  int
		d2  int
		res int = 0
	)
	if m == 0 || n == 0 {
		return res
	}
	for d1 = 1; d1 <= m; d1 *= 2 {
	}
	d1 /= 2
	for d2 = d1; d2 <= n; d2 *= 2 {
	}
	d2 /= 2
	// fmt.Printf("d1: %d d2: %d\n", d1, d2)
	// d1 <= d2
	if d1 == d2 {
		// end
		if d1 == 0 {
			return res
		}
		res += d1
		res += rangeBitwiseAnd(m-d1, n-d1)
	}
	return res
}

func main() {
	// test
	var (
		m   int = 999
		n   int = 1000
		res     = rangeBitwiseAnd(m, n)
	)
	fmt.Printf("rangeBitwiseAnd(%d(%b), %d(%b)) = %d(%b)\n", m, m, n, n, res, res)
}
