package main

import (
	"fmt"
)

// n >=2
func integerBreak(n int) int {

	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	var res int = 1
	for n >= 3 {
		res *= 3
		n -= 3
	}
	switch n {
	case 1:
		res = res / 3 * 4
	case 2:
		res *= 2
	}
	return res
}

func main() {

	// test
	n := 10
	fmt.Printf("integerBreak(%d) = %d\n", n, integerBreak(n))
}
