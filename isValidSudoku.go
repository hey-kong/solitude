package main

// Leetcode 36. (medium)
func isValidSudoku(board [][]byte) bool {
	row, col, box := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := range row {
		row[i], col[i], box[i] = make([]bool, 9), make([]bool, 9), make([]bool, 9)
	}
	for i := range board {
		for j := range board {
			if board[i][j] == '.' {
				continue
			}
			num := int(board[i][j] - '1')
			n := i/3*3 + j/3
			if row[i][num] || col[j][num] || box[n][num] {
				return false
			}
			row[i][num], col[j][num], box[n][num] = true, true, true
		}
	}
	return true
}
