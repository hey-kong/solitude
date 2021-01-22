package main

// Leetcode 419. (medium)
func countBattleships(board [][]byte) (res int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' && (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
				res++
			}
		}
	}
	return
}
