package main

import (
	"fmt"
)

func bulbSwitch(n int) int {
	res := 0
	for i := 1; i*i <= n; i++ {
		res += 1
	}
	return res
}

// TIME_OUT
/*func bulbSwitch(n int) int {
	if n == 0 {
		return 0
	}
	// calc count of factors, if odd then the bulb is on, otherwise, off
	var (
		res int = 0
		k   int = 1
	)
	for k <= n {
		if countFactors(k)%2 == 1 {
			res += 1
		}
		k += 1
	}
	return res
}

func countFactors(n int) int {
	// TODO: it can be faster
	var (
		k   int = n - 1
		res int = 1
	)
	for k > 0 {
		if n%k == 0 {
			res += 1
		}
		k -= 1
	}
	return res
}*/

func main() {

	// test
	n := 99999
	fmt.Printf("bulbSwitch(%d) = %d\n", n, bulbSwitch(n))
}
