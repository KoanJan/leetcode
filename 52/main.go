package main

import (
	"fmt"
)

func totalNQueens(n int) int {
	if n < 0 {
		panic("Invalid input")
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var res int
	for i := 0; i < n; i++ {
		var (
			board [][]byte = emptyChessBoard(n)
			crash [][]bool = emptyCrashRecords(n)
		)
		setQueen(board, crash, 0, i, n)
		res += calcNQueens(board, crash, n-1, n)
	}
	return res
}

func calcNQueens(board [][]byte, crash [][]bool, remaining, n int) int {
	if remaining == 0 {
		return 1
	}
	var (
		x, y, res int = -1, -1, 0
	)
	i := n - remaining
	for j := 0; j < n; j++ {
		if crash[i][j] {
			continue
		}
		var (
			newBoard [][]byte = cpBoard(board, n)
			newCrash [][]bool = cpCrash(crash, n)
		)
		x, y = i, j
		setQueen(newBoard, newCrash, x, y, n)
		res += calcNQueens(newBoard, newCrash, remaining-1, n)
	}
	return res
}

func setQueen(board [][]byte, crash [][]bool, x, y, n int) {
	board[x][y] = 'Q'
	for i1 := 0; i1 < n; i1++ {
		crash[i1][y] = true
		crash[x][i1] = true
	}
	for i1, j1 := x, y; i1 < n && j1 < n; {
		crash[i1][j1] = true
		i1++
		j1++
	}
	for i1, j1 := x, y; i1 >= 0 && j1 >= 0; {
		crash[i1][j1] = true
		i1--
		j1--
	}
	for i1, j1 := x, y; i1 >= 0 && j1 < n; {
		crash[i1][j1] = true
		i1--
		j1++
	}
	for i1, j1 := x, y; i1 < n && j1 >= 0; {
		crash[i1][j1] = true
		i1++
		j1--
	}
}

func emptyChessBoard(n int) [][]byte {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		board[i] = row
	}
	return board
}

func emptyCrashRecords(n int) [][]bool {
	crash := make([][]bool, n)
	for i := 0; i < n; i++ {
		crash[i] = make([]bool, n)
	}
	return crash
}

func cpBoard(board [][]byte, n int) [][]byte {
	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = board[i][j]
		}
		res[i] = row
	}
	return res
}

func cpCrash(crash [][]bool, n int) [][]bool {
	res := make([][]bool, n)
	for i := 0; i < n; i++ {
		row := make([]bool, n)
		for j := 0; j < n; j++ {
			row[j] = crash[i][j]
		}
		res[i] = row
	}
	return res
}

func main() {
	// test
	n := 9
	fmt.Printf("n = %d\ntotalNQueens: %v\n", n, totalNQueens(n))
}
