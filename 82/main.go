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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var (
		res *ListNode = new(ListNode)
		cur *ListNode = res
		dup int       = head.Val
	)
	// head
	if head.Val != head.Next.Val {
		cur.Next = &ListNode{Val: head.Val}
		cur = cur.Next
		head = head.Next
	}
	// loop
	for head.Next != nil {
		if head.Val != head.Next.Val && head.Val != dup {
			cur.Next = &ListNode{Val: head.Val}
			cur = cur.Next
		}
		dup = head.Val
		head = head.Next
	}
	// tail
	if head.Val != dup {
		cur.Next = &ListNode{Val: head.Val}
	}
	return res.Next
}

func main() {
	// test
	head := newListNode([]int{1, 1, 1, 2, 3})
	fmt.Printf("head: %v\n", head)
	head = deleteDuplicates(head)
	fmt.Printf("head without duplicates: %v\n", head)
}
