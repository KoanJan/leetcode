package main

import (
	"fmt"
)

// TODO 优化性能

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	b := []byte(s)
	rb := reverse(b)
	var max, cur, bestStart, bestEnd int
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			cur = equalLen(b[i:], rb[j:])
			if cur > max {
				max = cur
				bestStart = i
				bestEnd = i + cur
			}
		}
	}
	return string(b[bestStart:bestEnd])
}

func equalLen(b1, b2 []byte) int {
	ret := 0
	for i := 0; i < len(b1) && i < len(b2); i++ {
		if b1[i] != b2[i] {
			if i > 0 && b1[i-1] != b1[0]{
					return 0
			}
			return ret
		}
		ret++
	}
	if b1[ret-1] != b1[0] {
		return 0
	}
	return ret
}

func reverse(b []byte) []byte {
	if len(b) < 2 {
		return b
	}
	ret := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		ret[i] = b[len(b)-1-i]
	}
	return ret
}

func main() {
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
}
