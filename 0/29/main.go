package main

import (
	"fmt"
)

func divide(dividend int, divisor int) int {
	var (
		isMinus bool = false
		res     int
	)
	if dividend < 0 {
		dividend = -dividend
		isMinus = !isMinus
	}
	if divisor < 0 {
		divisor = -divisor
		isMinus = !isMinus
	}
	res = d(dividend, divisor)
	if isMinus {
		res = -res
	}
	// forking overflow
	if res > 2147483647 {
		res = 2147483647
	}
	return res
}

func d(dividend int, divisor int) int {
	if dividend < divisor {
		return 0
	}
	var (
		res, tmpDivisor int = 1, divisor
	)
	for dividend>>1 > tmpDivisor {
		tmpDivisor += tmpDivisor
		res += res
	}
	return res + d(dividend-tmpDivisor, divisor)
}

func main() {
	// test
	var dividend, divisor int = 55, -10
	fmt.Printf("%d / %d = %d\n", dividend, divisor, divide(dividend, divisor))
}
