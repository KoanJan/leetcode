package main

import "fmt"

func simplifyPath(path string) string {
	data := []byte(path)
	ln := len(data)
	i := 0
	stack := newSimpleStack()
	for i < ln {
		if data[i] == '/' {
			i++
			continue
		}
		if data[i] == '.' {
			if i+1 == ln || data[i+1] == '/' {
				// current dir
				i++
				continue
			} else if data[i+1] == '.' {
				// parent dir
				if i+2 == ln || data[i+2] == '/' {
					i += 2
					stack.Pop()
					continue
				}
			}
		}
		// hidden file
		path := []byte{}
		for ; i < ln && data[i] != '/'; i++ {
			path = append(path, data[i])
		}
		stack.Push(path)
	}
	// collect
	res := []byte{}
	for _, node := range stack.Data {
		res = append(res, '/')
		res = append(res, node...)
	}
	if len(res) == 0 {
		res = append(res, '/')
	}
	return string(res)
}

func newSimpleStack() *simpleStack {
	stack := new(simpleStack)
	stack.Data = [][]byte{}
	stack.ln = 0
	return stack
}

type simpleStack struct {
	Data [][]byte
	ln   int
}

func (stack *simpleStack) Push(name []byte) {
	stack.Data = append(stack.Data, name)
	stack.ln += 1
}

func (stack *simpleStack) Pop() {
	if stack.ln > 0 {
		stack.Data = stack.Data[0 : len(stack.Data)-1]
		stack.ln -= 1
	}
}

func main() {
	path := "/home/"
	fmt.Println(simplifyPath(path))
}
