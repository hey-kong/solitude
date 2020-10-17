package main

// Leetcode 52. (hard)
var cntOfTotalNQueens int

func totalNQueens(n int) int {
	cntOfTotalNQueens = 0
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	dfsTotalNQueens(n, 0, arr)
	return cntOfTotalNQueens
}

func dfsTotalNQueens(n, row int, arr [][]int) {
	if row == n {
		cntOfTotalNQueens++
		return
	}

	for i := 0; i < n; i++ {
		if isValidOfNQueens(n, row, i, arr) {
			arr[row][i] = 1
			dfsTotalNQueens(n, row+1, arr)
			arr[row][i] = 0
		}
	}
}

func isValidOfNQueens(n, row, col int, arr [][]int) bool {
	for i := 0; i < row; i++ {
		if arr[i][col] == 1 {
			return false
		}
	}
	i, j := row, col
	for i > 0 && j > 0 {
		i--
		j--
		if arr[i][j] == 1 {
			return false
		}
	}
	i, j = row, col
	for i > 0 && j < n-1 {
		i--
		j++
		if arr[i][j] == 1 {
			return false
		}
	}
	return true
}
