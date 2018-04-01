package main

import (
	"fmt"
)

type slidor struct {
	Val  int
	Next *slidor
}

func (this *slidor) Update(ov, nv int) {
	if ov == nv {
		return
	}
	cur := this
	for cur.Next != nil {
		if cur.Next.Val == ov {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	cur = this
	for cur.Next != nil {
		if cur.Next.Val <= nv {
			cur.Next = &slidor{Val: nv, Next: cur.Next}
			return
		}
		cur = cur.Next
	}
	cur.Next = &slidor{Val: nv}
}

func (this *slidor) Max() int {
	if this.Next != nil {
		return this.Next.Val
	}
	panic("empty slidor")
}

func (this *slidor) String() string {
	var (
		res string  = ""
		cur *slidor = this.Next
	)
	if cur == nil {
		return res
	} else {
		res = fmt.Sprintf("%d", cur.Val)
		cur = cur.Next
	}
	for cur != nil {
		res = fmt.Sprintf("%s %d", res, cur.Val)
	}
	return res
}

func NewSlidor(data []int) *slidor {
	this := new(slidor)
	if len(data) == 0 {
		return this
	}
	this.Next = &slidor{Val: data[0]}
	for i := 1; i < len(data); i++ {
		cur := this
		for cur.Next != nil {
			if cur.Next.Val <= data[i] {
				cur.Next = &slidor{Val: data[i], Next: cur.Next}
				break
			}
			cur = cur.Next
		}
		if cur.Next == nil {
			cur.Next = &slidor{Val: data[i]}
		}
	}
	return this
}

func maxSlidingWindow(nums []int, k int) []int {
	if k <= 0 || k > len(nums) || len(nums) == 0 {
		return []int{}
	}
	var (
		s   *slidor = NewSlidor(nums[0:k])
		res []int   = make([]int, len(nums)+1-k)
	)
	res[0] = s.Max()
	for i := k; i < len(nums); i++ {
		s.Update(nums[i-k], nums[i])
		res[i-k+1] = s.Max()
	}
	return res
}

func main() {
	// test
	var (
		nums []int = []int{1, 3, -1, -3, 5, 3, 6, 7}
		k    int   = 3
	)
	fmt.Printf("nums: %v\nk: %d\nresult: %v\n", nums, k, maxSlidingWindow(nums, k))
}
