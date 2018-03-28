package main

import "fmt"

func isValid(s string) bool {
	data := []byte(s)
	stack := newStack()
	for _, b := range data {
		switch b {
		case '}':
			if stack.empty() || stack.end() != '{' {
				return false
			}
			stack.pop()
		case ']':
			if stack.empty() || stack.end() != '[' {
				return false
			}
			stack.pop()
		case ')':
			if stack.empty() || stack.end() != '(' {
				return false
			}
			stack.pop()
		default:
			stack.push(b)
		}
	}
	return stack.empty()
}

type stack struct {
	data []byte
}

func newStack() *stack {
	s := new(stack)
	s.data = []byte{}
	return s
}

func (s *stack) empty() bool {
	return len(s.data) == 0
}

func (s *stack) pop() byte {
	i := len(s.data) - 1
	b := s.data[i]
	s.data = s.data[:i]
	return b
}

func (s *stack) end() byte {
	return s.data[len(s.data)-1]
}

func (s *stack) push(v byte) {
	s.data = append(s.data, v)
}

func main() {
	s := "{[()]}{{}()}"
	fmt.Println(isValid(s))
}
