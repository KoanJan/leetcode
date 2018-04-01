package main

import "fmt"

func longestValidParentheses(s string) int {
	data := []byte(s)
	// filt
	fmt.Println("filt")
	for len(data) > 0 {
		if data[0] == ')' {
			data = data[1:]
		} else if data[len(data)-1] == '(' {
			data = data[:len(data)-1]
		} else {
			break
		}
	}
	ln := len(data)
	if ln < 2 {
		return 0
	}
	// there are ln+1 node and ln steps
	fmt.Println("rec")
	rec := initArray(ln+1, ln+1, ln+1)
	for i := 0; i < ln+1; i++ {
		rec[i][i] = 0
	}
	for i := range data {
		if data[i] == '(' {
			rec[i][i+1] = 1
		} else {
			rec[i][i+1] = -1
		}
	}
	// DP
	fmt.Println("DP")
	for i := 0; i < ln-1; i++ {
		for j := 2; j < ln+1; j++ {
			for k := i + 1; k <= j; k++ {
				rec[i][k] = rec[i][k-1] + rec[k-1][k]
			}
		}
	}
	printArray(rec)
	// find
	fmt.Println("find")
	for length := ln; length > 0; length-- {
		for i := 0; i+length < ln+1; i++ {
			if rec[i][i+length] == 0 {
				if rec[i][i+1] == -1 {
					continue
				}
				return length
			}
		}
	}
	return 0
}

func initArray(x, y, init int) [][]int {
	a := make([][]int, x)
	for i := range a {
		subA := make([]int, y)
		for j := range subA {
			subA[j] = init
		}
		a[i] = subA
	}
	return a
}

func printArray(a [][]int) {
	for i := range a {
		for j := range a[i] {
			if j >= i {
				fmt.Printf("%d ", a[i][j])
			} else {
				fmt.Printf("_ ")
			}
		}
		fmt.Println("")
	}
}

func main() {
	s := "())()"
	fmt.Println(longestValidParentheses(s))
}
