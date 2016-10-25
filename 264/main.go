package main

import (
	"fmt"
)

func nthUglyNumber(n int) int {
	var (
		data       []int = []int{1, 2, 3, 4, 5}
		cur        int   = 6
		l1, l2, l3       = 1, 1, 1 // for calc min(l1*2, l2*3, l3*5)
	)
	if n <= 5 {
		return data[n-1]
	}
	for cur <= n {
		// l1, l2, l3
		for data[l1]*2 <= data[cur-2] {
			l1++
		}
		for data[l2]*3 <= data[cur-2] {
			l2++
		}
		for data[l3]*5 <= data[cur-2] {
			l3++
		}
		// calc min
		tmp := data[l1] * 2
		if data[l2]*3 < tmp {
			tmp = data[l2] * 3
		}
		if data[l3]*5 < tmp {
			tmp = data[l3] * 5
		}
		data = append(data, tmp)
		cur++
	}
	return data[cur-2]
}

func main() {
	// test
	n := 999
	fmt.Printf("The %dth UglyNumber is %d\n", n, nthUglyNumber(n))
}
