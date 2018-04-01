package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	var (
		l   int = 1
		r   int = 2147483647
		mid     = l/2 + r/2
	)
	if l%2 == 1 && r%2 == 1 {
		mid++
	}
	for r-l > 1 {
		tmp := mid * mid
		if tmp == x {
			return mid
		}
		if tmp < x {
			l = mid
		} else {
			r = mid
		}
		mid = l/2 + r/2
		if l%2 == 1 && r%2 == 1 {
			mid++
		}
	}
	tmp := r * r
	if tmp < x {
		return -2147483648
	} else if tmp == x {
		return r
	} else {
		return mid
	}
}

func main() {
	// test
	var x int = 10000000000000000
	fmt.Printf("mySqrt(%d)=%d\n", x, mySqrt(x))
}
