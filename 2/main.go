package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var (
		res   *ListNode = new(ListNode)
		cur             = res
		addon int       = 0
	)
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + addon
		cur.Next = &ListNode{Val: val % 10}
		addon = val / 10
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		val := l1.Val + addon
		cur.Next = &ListNode{Val: val % 10}
		addon = val / 10
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		val := l2.Val + addon
		cur.Next = &ListNode{Val: val % 10}
		addon = val / 10
		cur = cur.Next
		l2 = l2.Next
	}
	if addon != 0 {
		cur.Next = &ListNode{Val: addon}
	}
	return res.Next
}

func main() {
	// test
	var (
		l1 *ListNode = newListNode([]int{9, 1, 6})
		l2 *ListNode = newListNode([]int{0})
	)
	fmt.Printf("\t%v\n +      %v\n\n =      %v\n", l1, l2, addTwoNumbers(l1, l2))
}
