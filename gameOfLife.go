package main

// Leetcode 289. (medium)
func gameOfLife(board [][]int) {
	n, m := len(board), len(board[0])
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = board[i][j]
		}
	}

	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cnt := 0
			for k := 0; k < 8; k++ {
				x, y := i+dx[k], j+dy[k]
				if x >= 0 && x < n && y >= 0 && y < m && memo[x][y] == 1 {
					cnt++
				}
			}
			if cnt < 2 || cnt > 3 {
				board[i][j] = 0
			}
			if cnt == 3 && board[i][j] == 0 {
				board[i][j] = 1
			}
		}
	}
}
