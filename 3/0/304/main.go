package main

import (
	"fmt"
)

type NumMatrix struct {
	matrix [][]int
	sum    [][]int
}

func Construct(matrix [][]int) NumMatrix {
	numMatrix := NumMatrix{matrix: matrix}
	m := len(matrix)
	if m == 0 {
		return numMatrix
	}
	n := len(matrix[0])
	if n == 0 {
		return numMatrix
	}
	numMatrix.sum = make([][]int, m)
	for i := 0; i < m; i++ {
		numMatrix.sum[i] = make([]int, n)
	}
	numMatrix.sum[0][0] = matrix[0][0]
	for i := 1; i < m; i++ {
		numMatrix.sum[i][0] = numMatrix.sum[i-1][0] + matrix[i][0]
	}
	for i := 1; i < n; i++ {
		numMatrix.sum[0][i] = numMatrix.sum[0][i-1] + matrix[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			numMatrix.sum[i][j] = numMatrix.sum[i-1][j] + numMatrix.sum[i][j-1] - numMatrix.sum[i-1][j-1] + matrix[i][j]
		}
	}
	return numMatrix
}

func (this *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	fmt.Printf("this.matrix: %v\nthis.sum: %v\n", this.matrix, this.sum)
	res := this.sum[row2][col2]
	if row1 > 0 {
		res -= this.sum[row1-1][col2]
	}
	if col1 > 0 {
		res -= this.sum[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		res += this.sum[row1-1][col1-1]
	}
	return res
}

func main() {
	// test
	var (
		numMatrix NumMatrix = Construct([][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		})
		row1, col1, row2, col2 int
	)
	row1, col1, row2, col2 = 0, 0, 2, 2
	fmt.Printf("sumRegion(%d, %d, %d, %d) = %d\n", row1, col1, row2, col2, numMatrix.SumRegion(row1, col1, row2, col2))
	row1, col1, row2, col2 = 1, 1, 2, 2
	fmt.Printf("sumRegion(%d, %d, %d, %d) = %d\n", row1, col1, row2, col2, numMatrix.SumRegion(row1, col1, row2, col2))
}

///////////////////////////
// AC-ed JavaScript code //
///////////////////////////

// /**
//  * @constructor
//  * @param {number[][]} matrix
//  */
// var NumMatrix = function(matrix) {
//     this.matrix = matrix;
//     this.sum = [];
//     var m = matrix.length;
//     if (m === 0){
//         return;
//     }
//     var n = matrix[0].length;
//     if (n === 0){
//         return;
//     }
//     var i, j, tmp;
//     for (i = 0; i < m; i++){
//         tmp = [];
//         for (j = 0; j < n; j++){
//             tmp.push(0);
//         }
//         this.sum.push(tmp);
//     }
//     this.sum[0][0] = matrix[0][0];
//     for (i = 1; i < m; i++){
//         this.sum[i][0] = this.sum[i-1][0] + matrix[i][0];
//     }
//     for (i = 1; i < n; i++){
//         this.sum[0][i] = this.sum[0][i-1] + matrix[0][i];
//     }
//     for (i = 1; i < m; i++){
//         for (j = 1; j < n; j++){
//             this.sum[i][j] = this.sum[i][j-1] + this.sum[i-1][j] - this.sum[i-1][j-1] + this.matrix[i][j];
//         }
//     }
// };

// /**
//  * @param {number} row1
//  * @param {number} col1
//  * @param {number} row2
//  * @param {number} col2
//  * @return {number}
//  */
// NumMatrix.prototype.sumRegion = function(row1, col1, row2, col2) {
//     var res = this.sum[row2][col2];
//     if (row1 > 0){
//         res -= this.sum[row1-1][col2];
//     }
//     if (col1 > 0){
//         res -= this.sum[row2][col1-1];
//     }
//     if (row1 > 0 && col1 > 0){
//         res += this.sum[row1-1][col1-1];
//     }
//     return res;
// };

// /**
//  * Your NumMatrix object will be instantiated and called as such:
//  * var numMatrix = new NumMatrix(matrix);
//  * numMatrix.sumRegion(0, 1, 2, 3);
//  * numMatrix.sumRegion(1, 2, 3, 4);
//  */
