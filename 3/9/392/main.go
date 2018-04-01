package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	var (
		bs   []byte = []byte(s)
		bt   []byte = []byte(t)
		lenS int    = len(bs)
		lenT int    = len(bt)
		i    int    = 0
		j    int    = 0
	)
	for i < lenS && j < lenT {
		if bs[i] == bt[j] {
			i += 1
		}
		j += 1
	}
	return i == lenS
}

func main() {

	// test
	var (
		s string = "abdcd"
		t string = "aarhubhcohrdrge"
	)
	fmt.Printf("isSubsequence(%s, %s) = %v", s, t, isSubsequence(s, t))
}
