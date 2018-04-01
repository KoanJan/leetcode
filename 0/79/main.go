package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}
	if len([]byte(word)) > m*n {
		return false
	}
	passed := make([][]bool, m)
	for i := 0; i < m; i++ {
		passed[i] = make([]bool, n)
	}
	data := []byte(word)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == data[0] {
				tmp := cpMatrix(passed)
				tmp[i][j] = true
				if e(board, data[1:], tmp, i, j, m, n) {
					return true
				}
			}
		}
	}
	return false
}

func e(board [][]byte, data []byte, passed [][]bool, x, y, m, n int) bool {
	if len(data) == 0 {
		return true
	}
	if x > 0 {
		if !passed[x-1][y] && board[x-1][y] == data[0] {
			tmp := cpMatrix(passed)
			tmp[x-1][y] = true
			if e(board, data[1:], tmp, x-1, y, m, n) {
				return true
			}
		}
	}
	if x < m-1 {
		if !passed[x+1][y] && board[x+1][y] == data[0] {
			tmp := cpMatrix(passed)
			tmp[x+1][y] = true
			if e(board, data[1:], tmp, x+1, y, m, n) {
				return true
			}
		}
	}
	if y > 0 {
		if !passed[x][y-1] && board[x][y-1] == data[0] {
			tmp := cpMatrix(passed)
			tmp[x][y-1] = true
			if e(board, data[1:], tmp, x, y-1, m, n) {
				return true
			}
		}
	}
	if y < n-1 {
		if !passed[x][y+1] && board[x][y+1] == data[0] {
			tmp := cpMatrix(passed)
			tmp[x][y+1] = true
			if e(board, data[1:], tmp, x, y+1, m, n) {
				return true
			}
		}
	}
	return false
}

func cpMatrix(src [][]bool) [][]bool {
	res := [][]bool{}
	for i := 0; i < len(src); i++ {
		tmp := []bool{}
		for j := 0; j < len(src[i]); j++ {
			tmp = append(tmp, src[i][j])
		}
		res = append(res, tmp)
	}
	return res
}

func main() {
	// test
	var (
		board [][]byte = [][]byte{
			[]byte{'A', 'B', 'C', 'E'},
			[]byte{'S', 'F', 'C', 'S'},
			[]byte{'A', 'D', 'E', 'E'},
		}
		word string = "ABCB"
	)
	fmt.Printf("\"%s\" in matrix:\n%v\n\n%t\n", word, board, exist(board, word))
}
