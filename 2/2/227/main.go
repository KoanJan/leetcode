package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type calcListNode struct {
 *     Val int
 *     Next *calcListNode
 * }
 */
type calcListNode struct {
	Val      []byte
	IsNumber bool
	Next     *calcListNode
}

func (c *calcListNode) String() string {
	if c == nil {
		return "<nil>"
	}
	var (
		res string
		cur *calcListNode = c
	)
	for cur != nil {
		var vals string
		for i := 0; i < len(cur.Val); i++ {
			vals += fmt.Sprintf("%c", cur.Val[i])
		}
		res += vals
		if cur.Next != nil {
			res += " -> "
		}
		cur = cur.Next
	}
	return res
}

func calculate(s string) int {
	var (
		formula   *calcListNode = &calcListNode{IsNumber: false}
		bs        []byte        = []byte(s)
		cv        []byte        = []byte{}
		cur                     = formula
		calc      *calcListNode
		curNumber []int  = []int{}
		curSymbol []byte = []byte{}
	)
	if len(bs) == 0 {
		return 0
	}

	// store as linkedlist
	for i := 0; i < len(bs); i++ {
		if bs[i] == ' ' {
			continue
		}
		if isNumber(bs[i]) {
			cv = append(cv, bs[i])
		} else {
			cur.Next = &calcListNode{Val: cv, IsNumber: true}
			cur = cur.Next
			cv = []byte{}
			cur.Next = &calcListNode{Val: []byte{bs[i]}, IsNumber: false}
			cur = cur.Next
		}
	}
	cur.Next = &calcListNode{Val: cv, IsNumber: true}
	cur = formula.Next

	// fmt.Printf("formula: %v\n", formula)

	// 1. calc '*' and '/'
	for cur != nil {
		if cur.IsNumber {
			// calc
			if len(curNumber) == 0 {
				curNumber = append(curNumber, cvtNumber(cur.Val))
				//
				calc = cur
			} else if len(curSymbol) != 0 {
				if curSymbol[0] == '*' {
					curNumber[0] = curNumber[0] * cvtNumber(cur.Val)
				} else if curSymbol[0] == '/' {
					curNumber[0] = curNumber[0] / cvtNumber(cur.Val)
				} else {
					panic("Unknown wrong!")
				}
				curSymbol = []byte{}
			} else {
				panic("Invalid input!")
			}
		} else {
			// end unity calc and short formula
			// store symbol
			if cur.Val[0] == '*' || cur.Val[0] == '/' {
				curSymbol = append(curSymbol, cur.Val[0])
			} else {
				// update vals and short formula
				calc.Val = cvtBts(curNumber[0])
				calc.Next = cur
				// reset vals
				calc = nil
				curNumber = []int{}
				curSymbol = []byte{}
			}
		}
		cur = cur.Next
	}
	// if tail
	if len(curNumber) != 0 {
		// update vals and short formula
		calc.Val = cvtBts(curNumber[0])
		calc.Next = nil
		// reset vals
		calc = nil
		curNumber = []int{}
		curSymbol = []byte{}
	}

	// 2. calc '+' and '-'
	// fmt.Printf("formula: %v\n", formula)
	cur = formula.Next
	curNumber = append(curNumber, 0)
	curSymbol = append(curSymbol, '+')
	for cur != nil {
		if cur.IsNumber {
			if len(curSymbol) != 0 {
				if curSymbol[0] == '+' {
					curNumber[0] += cvtNumber(cur.Val)
				} else if curSymbol[0] == '-' {
					curNumber[0] -= cvtNumber(cur.Val)
				} else {
					panic("Invalid input!")
				}
				curSymbol = []byte{}
			} else {
				panic("Unknown wrong!")
			}
		} else {
			curSymbol = append(curSymbol, cur.Val[0])
		}
		cur = cur.Next
	}

	return curNumber[0]
}

func isNumber(c byte) bool {
	return 48 <= c && c <= 57
}

func cvtNumber(bs []byte) int {
	var (
		res int = 0
		d   int = 1
	)
	for i := len(bs) - 1; i >= 0; i-- {
		res += (int(bs[i]) - 48) * d
		d *= 10
	}
	return res
}

func cvtBts(n int) []byte {
	if n == 0 {
		return []byte{'0'}
	}

	var (
		res []byte = []byte{}
		tmp []byte = []byte{}
	)
	for n > 0 {
		tmp = append(tmp, byte(n%10+48))
		n /= 10
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}
	return res
}

func main() {
	// test
	s := "0/1"
	fmt.Printf("%s=%d\n", s, calculate(s))
}
