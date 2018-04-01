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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}

	var (
		res    *ListNode = new(ListNode)
		rev    *ListNode = new(ListNode)
		cur    *ListNode = head
		resCur *ListNode = res
	)
	for i := 1; i < m; i++ {
		resCur.Next = &ListNode{Val: cur.Val}
		resCur = resCur.Next
		cur = cur.Next
	}
	for i := m; i <= n; i++ {
		rev.Next = &ListNode{Val: cur.Val, Next: rev.Next}
		cur = cur.Next
	}
	resCur = res
	for resCur.Next != nil {
		resCur = resCur.Next
	}
	if rev.Next != nil {
		resCur.Next = rev.Next
	}
	for resCur.Next != nil {
		resCur = resCur.Next
	}
	resCur.Next = cur
	return res.Next
}

func main() {
	// test
	var (
		head *ListNode = newListNode([]int{1, 2})
		m    int       = 1
		n    int       = 2
	)
	fmt.Printf("head: %v\nm: %d\nn: %d\n\n", head, m, n)
	head = reverseBetween(head, m, n)
	fmt.Printf("reversed: %v\n", head)
}
