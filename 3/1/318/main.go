package main

import (
	"fmt"
)

func maxProduct(words []string) int {

	// init
	var (
		max  int   = 0
		size int   = len(words)
		nums []int = make([]int, size)
	)
	for i := 0; i < size; i++ {
		tmp := words[i]
		for j := 0; j < stringLength(tmp); j++ {
			nums[i] |= (1 << (charAt(tmp, j) - 'a'))
		}
	}

	for i := 0; i < size; i++ {
		for j := size - 1; j > i; j-- {
			if nums[i]&nums[j] == 0 {
				product := stringLength(words[i]) * stringLength(words[j])
				if product > max {
					max = product
				}
			}
		}
	}

	return max

}

func stringLength(str string) int {
	return len([]byte(str))
}

func charAt(str string, idx int) byte {
	return []byte(str)[idx]
}

func main() {

	// test
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Printf("maxProduct(words) = %d while words is %v\n", maxProduct(words), words)
}
