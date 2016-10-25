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

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	var (
		left      *ListNode = &ListNode{}
		leftTail  *ListNode = left
		right     *ListNode = &ListNode{}
		rightTail *ListNode = right
		cur       *ListNode = head
		res       *ListNode = &ListNode{}
	)
	// partition
	for cur != nil {
		if cur.Val < x {
			leftTail.Next = &ListNode{Val: cur.Val}
			leftTail = leftTail.Next
		} else {
			rightTail.Next = &ListNode{Val: cur.Val}
			rightTail = rightTail.Next
		}
		cur = cur.Next
	}
	// link
	res.Next = left.Next
	if res.Next != nil {
		leftTail.Next = right.Next
	} else {
		res.Next = right.Next
	}
	return res.Next
}

func main() {
	// test
	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	x := 4

	fmt.Printf("head: %v\n", head)
	head = partition(head, x)
	fmt.Printf("after partition, head: %v\n", head)
}
