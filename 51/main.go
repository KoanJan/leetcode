package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	var res [][]string = [][]string{}
	if n < 0 {
		panic("Invalid input")
	}
	if n == 0 {
		return res
	}
	if n == 1 {
		res = append(res, []string{"Q"})
		return res
	}
	for i := 0; i < n; i++ {
		var (
			board [][]byte = emptyChessBoard(n)
			crash [][]bool = emptyCrashRecords(n)
		)
		setQueen(board, crash, 0, i, n)
		subRes := snq(board, crash, n-1, n)
		for j := 0; j < len(subRes); j++ {
			res = append(res, fmtOKSolution(subRes[j]))
		}
	}
	return res
}

func snq(board [][]byte, crash [][]bool, remaining, n int) [][][]byte {
	if remaining == 0 {
		return [][][]byte{board}
	}
	var (
		x, y int        = -1, -1
		res  [][][]byte = [][][]byte{}
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
		subRes := snq(newBoard, newCrash, remaining-1, n)
		for i1 := 0; i1 < len(subRes); i1++ {
			res = append(res, subRes[i1])
		}
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

func fmtOKSolution(solution [][]byte) []string {
	res := make([]string, len(solution))
	for i := 0; i < len(solution); i++ {
		res[i] = string(solution[i])
	}
	return res
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

//////////////////////
// output functions //
//////////////////////

func boardString(board [][]byte, n int) string {
	var res string
	for i := 0; i < n; i++ {
		var row string = fmt.Sprintf("%c", board[i][0])
		for j := 1; j < n; j++ {
			row = fmt.Sprintf("%s %c ", row, board[i][j])
		}
		res += row + "\n"
	}
	return res
}

func crashString(crash [][]bool, n int) string {
	var res string
	for i := 0; i < n; i++ {
		var row string = fmt.Sprintf("%t", crash[i][0])
		for j := 1; j < n; j++ {
			row = fmt.Sprintf("%s %t ", row, crash[i][j])
		}
		res += row + "\n"
	}
	return res
}

func fmtSolutions(solutions [][]string) string {
	var res string
	for i := 0; i < len(solutions); i++ {
		var row string
		for j := 0; j < len(solutions[i]); j++ {
			row += solutions[i][j] + "\n"
		}
		res += row + "\n"
	}
	return res
}

func main() {
	// test
	n := 4
	fmt.Printf("n = %d\nNQueens solution: %v\n", n, solveNQueens(n))
}
