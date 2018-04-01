package main

import (
	"fmt"
)

func validUtf8(data []int) bool {
	// init
	db := [][]int{}
	for i := 0; i < len(data); i++ {
		db = append(db, intToBin(data[i]))
	}

	// empty input
	if len(db) == 0 {
		return false
	}

	// calc bytes
	k := 0
	for k < len(db) {
		n := 0
		for i := 0; i < len(db[k]); i++ {
			if db[k][i] == 1 {
				n++
			} else {
				break
			}
		}
		if n > 1 {
			// n-bytes

			if len(db) < k+n {
				return false
			}
			for i := k + 1; i < k+n; i++ {
				if db[i][0] != 1 || db[i][1] != 0 {
					return false
				}
			}
			k = k + n
		} else if n == 0 {
			// 1 byte
			k++
		} else {
			return false
		}
	}
	return true
}

func intToBin(a int) []int {
	res := make([]int, 8)
	for i := len(res) - 1; i >= 0 && a > 0; i-- {
		res[i] = a % 2
		a /= 2
	}
	return res
}

func main() {
	// test
	data := []int{145}
	fmt.Printf("data = %v, validUtf8(data) = %t\n", data, validUtf8(data))
}
