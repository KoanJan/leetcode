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

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	var (
		length int       = 0
		cur    *ListNode = head
	)
	for cur != nil {
		length++
		cur = cur.Next
	}
	k %= length
	if k == 0 {
		return head
	}

	cur = head
	var newHead, newTail, oldTail *ListNode
	for i := 0; i < length-k-1; i++ {
		cur = cur.Next
	}
	newTail = cur
	newHead = cur.Next
	for cur.Next != nil {
		cur = cur.Next
	}
	oldTail = cur

	newTail.Next = nil
	oldTail.Next = head
	return newHead
}

func main() {
	// test
	var (
		head *ListNode = newListNode([]int{1})
		k    int       = 0
	)
	fmt.Printf("head: %v\nk: %d\n rotatedRight: %v\n", head, k, rotateRight(head, k))
}
