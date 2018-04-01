package main

import (
	"fmt"
)

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}
	if z > x+y {
		return false
	}
	return z%gcd(x, y) == 0
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	// test
	var x, y, z int = 3, 5, 4
	fmt.Printf("canMeasureWater(%d, %d, %d) -- %t\n", x, y, z, canMeasureWater(x, y, z))
}
