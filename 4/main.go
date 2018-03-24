package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ln := len(nums1) + len(nums2)
	k := ln/2 + 1
	if ln%2 == 0 {
		return (findMin(nums1, nums2, k-1) + findMin(nums1, nums2, k)) / 2.0
	} else {
		return findMin(nums1, nums2, k)
	}
}

func findMin(nums1 []int, nums2 []int, k int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}
	if m == 0 {
		return float64(nums2[k-1])
	}
	if k == 1 {
		return float64(min(nums1[0], nums2[0]))
	}
	pm := min(k/2, m)
	pn := k - pm
	// attempt that the top 'k' elements contain 'pm' elements from nums1 and 'pn' from nums2
	if nums1[pm-1] > nums2[pn-1] {
		return findMin(nums1, nums2[pn:], pm)
	} else {
		return findMin(nums1[pm:], nums2, pn)
	}

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	//nums1 := []int{1, 4, 5, 6, 7, 8, 9}
	//nums2 := []int{2, 3}
	nums1 := []int{1, 2, 6, 7, 8}
	nums2 := []int{3, 4, 5}
	m := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("%f", m)
}
