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

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	var (
		cur     *ListNode = head
		preTail *ListNode = head
		tail    *ListNode = head.Next
	)
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {

		// relink
		preTail = cur
		for preTail.Next.Next != nil {
			preTail = preTail.Next
		}
		tail = preTail.Next
		preTail.Next = nil
		tail.Next = cur.Next
		cur.Next = tail

		cur = cur.Next.Next
	}
}

func main() {
	// test
	head := newListNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("head: %v\n", head)
	reorderList(head)
	fmt.Printf("reordered: %v\n", head)
}
