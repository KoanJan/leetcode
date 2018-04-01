package main

import (
	"fmt"
)

func largestNumber(nums []int) string {
	var (
		byteNums [][]byte = [][]byte{}
		res      []byte   = []byte{}
		isZero   bool     = true
	)
	for i := 0; i < len(nums); i++ {
		byteNums = append(byteNums, int2BytesR(nums[i]))
	}
	byteNums = mergeSort(byteNums)
	for i := 0; i < len(byteNums); i++ {
		for j := len(byteNums[i]) - 1; j >= 0; j-- {
			if byteNums[i][j] != '0' {
				isZero = false
			}
			res = append(res, byteNums[i][j])
		}
	}
	if isZero {
		return "0"
	}
	return string(res)
}

func largerThan(bnum1, bnum2 []byte) bool {
	var i1, i2 int = len(bnum1) - 1, len(bnum2) - 1

	for i1 >= 0 && i2 >= 0 {
		if bnum1[i1] < bnum2[i2] {
			return false
		} else if bnum1[i1] > bnum2[i2] {
			return true
		}
		i1--
		i2--
	}
	if i1 == i2 {
		return false
	}
	if i1 > i2 {
		return largerThan(bnum1[0:i1+1], bnum2)
	}
	if i1 < i2 {
		return largerThan(bnum1, bnum2[0:i2+1])
	}
	// this code won't run
	return false
}

func int2BytesR(num int) []byte {
	if num == 0 {
		return []byte{'0'}
	}
	res := []byte{}
	for num > 0 {
		res = append(res, byte(num%10+48))
		num /= 10
	}
	return res
}

// mergeSort
func mergeSort(a [][]byte) [][]byte {
	if len(a) < 2 {
		return a
	}
	return merge(mergeSort(a[0:len(a)/2]), mergeSort(a[len(a)/2:]))
}

func merge(a [][]byte, b [][]byte) [][]byte {
	var (
		res [][]byte = [][]byte{}
		ca  int      = 0
		cb  int      = 0
	)
	for {
		// index out
		if ca == len(a) || cb == len(b) {
			break
		}
		if largerThan(a[ca], b[cb]) {
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
	nums := []int{0, 0, 0}
	fmt.Printf("nums: %v\nlargestNumber: %s\n", nums, largestNumber(nums))
}
