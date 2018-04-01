package main

import (
	"fmt"
)

type RPNNode struct {
	Val  string
	Next *RPNNode
}

func (r *RPNNode) String() string {
	if r == nil {
		return "<nil>"
	}
	var (
		res string
		cur *RPNNode = r
	)
	for cur != nil {
		res += cur.Val
		if cur.Next != nil {
			res += " -> "
		}
		cur = cur.Next
	}
	return res
}

func newRPNNode(tokens []string) *RPNNode {
	var (
		res = new(RPNNode)
		cur = res
	)
	for i := 0; i < len(tokens); i++ {
		cur.Next = &RPNNode{Val: tokens[i]}
		cur = cur.Next
	}
	return res.Next
}

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	if len(tokens) == 1 {
		return parseInteger(tokens[0])
	}
	if len(tokens)%2 == 0 {
		panic("Invalid RPN")
	}
	var (
		data *RPNNode = &RPNNode{Next: newRPNNode(tokens)}
		c             = data
	)
	for data.Next.Next != nil {
		c = data
		if c.Next.Next.Next == nil {
			panic("Invalid input")
		}
		for c.Next.Next.Next != nil {
			if isSymbol(c.Next.Next.Next.Val) && !isSymbol(c.Next.Next.Val) && !isSymbol(c.Next.Val) {
				c.Next = &RPNNode{
					Val:  calc(c.Next.Val, c.Next.Next.Val, c.Next.Next.Next.Val),
					Next: c.Next.Next.Next.Next,
				}
			}
			c = c.Next
			if c == nil || c.Next == nil || c.Next.Next == nil {
				break
			}
		}
	}
	return parseInteger(data.Next.Val)
}

func isSymbol(token string) bool {
	return token == "+" || token == "/" || token == "*" || token == "-"
}

func calc(x, y, symbol string) string {
	var (
		_x  int = parseInteger(x)
		_y  int = parseInteger(y)
		res string
	)
	switch symbol {
	case "*":
		res = parseString(_x * _y)
	case "/":
		res = parseString(_x / _y)
	case "+":
		res = parseString(_x + _y)
	case "-":
		res = parseString(_x - _y)
	}
	return res
}

func parseInteger(val string) int {
	var (
		data    []byte = []byte(val)
		res     int    = 0
		isMinus bool   = false
		d       int    = 1
	)

	for i := len(data) - 1; i >= 0; i-- {
		if data[i] == '-' {
			isMinus = true
			break
		}
		res += (int(data[i]) - 48) * d
		d *= 10
	}
	if isMinus {
		res = -res
	}
	return res
}

func parseString(val int) string {
	if val == 0 {
		return "0"
	}
	var (
		data    []byte
		d       int
		isMinus bool = val < 0
		tmp     int
		res     string
	)
	if isMinus {
		val = -val
	}
	tmp = val
	for tmp > 0 {
		d++
		tmp /= 10
	}
	data = make([]byte, d)
	for i := d - 1; i >= 0 && val > 0; i-- {
		data[i] = byte(val%10 + 48)
		val /= 10
	}
	res = string(data)
	if isMinus {
		res = "-" + res
	}
	return res
}

func main() {
	// test
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Printf("tokens: %v\nresult: %d\n", tokens, evalRPN(tokens))
}
