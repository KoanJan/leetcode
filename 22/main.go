package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	return linkedList2StringSlice(gp(n, n, ""))
}

func gp(l, r int, base string) *linkedList {

	var res *linkedList

	if r == 0 && l == 0 {
		res = &linkedList{val: base}
		res.tail = res
		res.tail = res
		return res
	}

	if l > 0 {
		res = link(res, gp(l-1, r, base+"("))
	}
	if r-l > 0 {
		res = link(res, gp(l, r-1, base+")"))
	}
	return res
}

func linkedList2StringSlice(l *linkedList) []string {
	res := []string{}
	if l == nil {
		return res
	}
	for l != nil {
		res = append(res, l.val)
		l = l.next
	}
	return res
}

type linkedList struct {
	val  string
	next *linkedList
	tail *linkedList
}

func link(lists ...*linkedList) *linkedList {
	if len(lists) == 0 {
		return nil
	}
	if lists[0] == nil {
		return link(lists[1:]...)
	}
	res := lists[0]
	for i := 1; i < len(lists); i++ {
		if lists[i] != nil {
			res.tail.next = lists[i]
			res.tail = lists[i].tail
		}
	}
	return res
}

func main() {

	// test
	n := 10
	fmt.Printf("generateParenthesis(%d):\n%v\n", n, generateParenthesis(n))
}
