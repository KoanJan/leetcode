package main

import (
	"fmt"
)

func countBits(num int) []int {
	//
	if num < 0 {
		return make([]int, 0)
	} else if num == 0 {
		return make([]int, 1)
	} else if num == 1 {
		return []int{0, 1}
	}

	// common solution
	res := make([]int, num+1)
	var i int = 1
	for i <= num {
		res[i] = 1
		cbs(res, i/2, i)
		i *= 2
	}
	if i/2 < num {
		var (
			k = i/2 + 1
			m = 1
		)
		for k <= num {
			res[k] = res[m] + 1
			k += 1
			m += 1
		}
	}
	return res
}

func cbs(raw []int, a, b int) {
	c := (a + b) / 2
	if c == a {
		if raw[b] == 0 {
			raw[b] = raw[a] + 1
		}
		return
	}
	raw[c] = raw[a] + 1
	cbs(raw, a, c)
	cbs(raw, c, b)

}

func main() {
	// test
	num := 50
	fmt.Printf("num = %d, and result is %v", num, countBits(num))
}
