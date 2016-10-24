package main

import (
	"fmt"
)

func hIndex(citations []int) int {
	citations = mergeSort(citations)
	n := len(citations)
	for i := 0; i < n; i++ {
		if citations[i] < i+1 {
			return i
		}
	}
	return n
}

// mergeSort
func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	return merge(mergeSort(a[0:len(a)/2]), mergeSort(a[len(a)/2:]))
}

func merge(a []int, b []int) []int {
	var (
		res []int = []int{}
		ca  int   = 0
		cb  int   = 0
	)
	for {
		// index out
		if ca == len(a) || cb == len(b) {
			break
		}
		if a[ca] > b[cb] {
			res = append(res, a[ca])
			ca++
		} else {
			res = append(res, b[cb])
			cb++
		}
	}
	for ca < len(a) {
		res = append(res, a[ca])
		ca++
	}
	for cb < len(b) {
		res = append(res, b[cb])
		cb++
	}
	return res
}

func main() {
	// test
	citations := []int{4, 4, 7, 8, 2, 7, 7, 7, 5}
	fmt.Printf("H-Index of %v is %d\n", citations, hIndex(citations))
}
