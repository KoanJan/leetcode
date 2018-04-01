package main

import (
	"fmt"
)

func longestSubstring(s string, k int) int {
	var (
		b         []byte       = []byte(s)
		data      map[byte]int = map[byte]int{}
		maxLength int          = 0
		start     int          = 0
		end       int          = 0
	)
	if len(b) == 0 {
		return maxLength
	}
	for i := 0; i < len(b); i++ {
		data[b[i]] += 1
	}
	for end < len(b) {
		if data[b[end]] >= k {
			end++
		} else {
			tmp := longestSubstring(string(b[start:end]), k)
			if tmp > maxLength {
				maxLength = tmp
			}
			start = end + 1
			end = start
		}
	}
	if start == 0 {
		if len(b) > maxLength {
			maxLength = len(b)
			return maxLength
		}
	}
	if end-start+1 >= k {
		tmp := longestSubstring(string(b[start:]), k)
		if tmp > maxLength {
			maxLength = tmp
		}
	}
	return maxLength
}

func main() {
	// test
	var (
		s string = "bbaaacbd"
		k int    = 3
	)
	fmt.Printf("The length of longest substring of %s with at least %d repeats is %d\n", s, k, longestSubstring(s, k))
}
