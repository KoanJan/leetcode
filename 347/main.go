package main

import (
	"fmt"
)

type elem struct {
	V int
	F int
}

func topKFrequent(nums []int, k int) []int {
	var (
		data  = map[int]int{}
		array = []elem{}
		res   = []int{}
	)
	// save all frequece
	for _, num := range nums {
		data[num] += 1
	}
	// init array
	for v, f := range data {
		array = append(array, elem{V: v, F: f})
	}
	// sort ps: O(n) <= nlogn
	array = mergeSort(array)
	// make result
	for i := 0; i < k; i++ {
		res = append(res, array[i].V)
	}
	return res

}

// mergeSort
func mergeSort(a []elem) []elem {
	if len(a) < 2 {
		return a
	}
	return merge(mergeSort(a[0:len(a)/2]), mergeSort(a[len(a)/2:]))
}

func merge(a []elem, b []elem) []elem {
	var (
		res []elem = []elem{}
		ca  int    = 0
		cb  int    = 0
	)
	for {
		// index out
		if ca == len(a) || cb == len(b) {
			break
		}
		if a[ca].F > b[cb].F {
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
	var (
		nums = []int{1, 1, 1, 2, 2, 3}
		k    = 2
	)
	fmt.Printf("topKFrequent(nums, k) = %v while nums is %v and k is %d", topKFrequent(nums, k), nums, k)
}
