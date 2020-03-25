package main

// Leetcode 999. (easy)
func numRookCaptures(board [][]byte) int {
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, -1, 0, 1}

	res := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				for k := 0; k < 4; k++ {
					x, y := i, j
					for {
						x += dx[k]
						y += dy[k]
						if x < 0 || x >= 8 || y < 0 || y >= 8 || board[x][y] == 'B' {
							break
						}
						if board[x][y] == 'p' {
							res++
							break
						}
					}
				}
				return res
			}
		}
	}
	return res
}
