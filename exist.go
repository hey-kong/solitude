package main

// Leetcode 79. (medium)
func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	if len(board) == 0 {
		return false
	}

	mr, mc := len(board), len(board[0])
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			if board[i][j] == word[0] && dfsExist(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfsExist(board [][]byte, r, c int, word string, i int) bool {
	mr, mc := len(board), len(board[0])
	if r == -1 || r == mr || c == -1 || c == mc || i == len(word) || board[r][c] != word[i] {
		return false
	}
	if i == len(word)-1 && board[r][c] == word[i] {
		return true
	}

	b := board[r][c]
	board[r][c] = '0'
	if dfsExist(board, r-1, c, word, i+1) {
		return true
	}
	if dfsExist(board, r, c+1, word, i+1) {
		return true
	}
	if dfsExist(board, r+1, c, word, i+1) {
		return true
	}
	if dfsExist(board, r, c-1, word, i+1) {
		return true
	}
	board[r][c] = b
	return false
}
