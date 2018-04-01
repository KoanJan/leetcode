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

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var (
		i   *ListNode = head.Next
		res *ListNode = &ListNode{Val: head.Val}
		cur           = res
	)
	for i != nil {
		// head
		if i.Val < res.Val {
			res = &ListNode{Val: i.Val, Next: res}
		} else {
			// body
			for cur.Next != nil {
				if i.Val < cur.Next.Val {
					cur.Next = &ListNode{Val: i.Val, Next: cur.Next}
					break
				}
				cur = cur.Next
			}
			// tail
			if cur.Next == nil {
				cur.Next = &ListNode{Val: i.Val}
			}
		}
		i = i.Next
		cur = res
	}

	return res
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

	fmt.Printf("head: %v\n", head)
	head = insertionSortList(head)
	fmt.Printf("sorted head: %v\n", head)
}
