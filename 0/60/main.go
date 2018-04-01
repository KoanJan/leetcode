package main

import "fmt"

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	kinds := factor(n)
	return permutation(nums, kinds, k)
}

func permutation(nums []int, kinds, k int) string {
	// fmt.Printf("param nums is %v and kinds is %d and k is %d\n", nums, kinds, k)
	ln := len(nums)
	if ln == 1 {
		return fmt.Sprintf("%d", nums[0])
	}
	newKinds := kinds / ln
	numIndex := (k - 1) / newKinds
	num := nums[numIndex]
	nums = append(nums[0:numIndex], nums[numIndex+1:]...)
	k %= newKinds
	if k == 0 {
		k = newKinds
	}
	return fmt.Sprintf("%d%s", num, permutation(nums, newKinds, k))
}

func factor(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	n, k := 3, 2
	fmt.Println(getPermutation(n, k))
}
