package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	if num < 1 {
		return false
	}
	if num == 1 {
		return true
	}
	var (
		l   int = 1
		r   int = num
		mid int = (l + r) / 2
	)
	for l <= r {
		tmp := mid * mid
		if tmp == num {
			return true
		} else if tmp < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
		mid = (l + r) / 2
	}
	return false
}

func main() {
	// test
	num := 16
	fmt.Printf("isPerfectSquare(%d): %t\n", num, isPerfectSquare(num))
}
