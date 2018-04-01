package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "<nil>"
	}
	var (
		res string
		cur *ListNode = l
	)
	for cur != nil {
		res += fmt.Sprintf("%d", cur.Val)
		if cur.Next != nil {
			res += " -> "
		}
		cur = cur.Next
	}
	return res
}

func newListNode(nums []int) *ListNode {
	var (
		res = new(ListNode)
		cur = res
	)
	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return res.Next
}

func multiply(num1 string, num2 string) string {
	if num1 == "" || num2 == "" {
		panic("Invalid input")
	}

	// init
	var (
		b1    []byte    = []byte(num1)
		b2    []byte    = []byte(num2)
		l1    *ListNode = new(ListNode)
		l2    *ListNode = new(ListNode)
		data  *ListNode = newListNode(make([]int, len(b1)+len(b2)))
		c1              = l1
		c2              = l2
		c               = data
		addon int       = 0
		res   []byte    = make([]byte, len(b1)+len(b2))
	)
	for i := len(b1) - 1; i >= 0; i-- {
		c1.Next = &ListNode{Val: int(b1[i]) - 48}
		c1 = c1.Next
	}
	c1 = l1.Next
	for i := len(b2) - 1; i >= 0; i-- {
		c2.Next = &ListNode{Val: int(b2[i] - 48)}
		c2 = c2.Next
	}
	c2 = l2.Next

	// calc
	times := c
	for c2 != nil {
		c = times
		c1 = l1.Next
		for c1 != nil {
			all := c.Val + c1.Val*c2.Val + addon
			c.Val = all % 10
			addon = all / 10
			c1 = c1.Next
			c = c.Next
		}
		if addon != 0 {
			c.Val = addon
			addon = 0
		}
		times = times.Next
		c2 = c2.Next
	}

	c = data
	// result
	for i := len(res) - 1; i >= 0 && c != nil; i-- {
		res[i] = byte(c.Val + 48)
		c = c.Next
	}
	for i := 0; i < len(res); i++ {
		if res[i] != '0' {
			// cut zeros
			return string(res[i:])
		}
	}
	return "0"
}

func main() {
	// test
	var (
		num1 string = "9"
		num2 string = "99"
	)
	fmt.Printf("%s * %s = %s\n", num1, num2, multiply(num1, num2))
}
