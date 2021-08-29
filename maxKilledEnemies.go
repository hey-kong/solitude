package main

// Leetcode 361. (medium)
func maxKilledEnemies(grid [][]byte) int {
	row, col := len(grid), len(grid[0])
	dp := make([][]int, row)
	for i := range grid {
		dp[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		left, right := 0, 0
		for j := 0; j < col; j++ {
			rc := col - 1 - j
			dp[i][j] += left
			dp[i][rc] += right
			if grid[i][j] == 'W' {
				left = 0
			} else if grid[i][j] == 'E' {
				left++
			}
			if grid[i][rc] == 'W' {
				right = 0
			} else if grid[i][rc] == 'E' {
				right++
			}
		}
	}
	for j := 0; j < col; j++ {
		up, down := 0, 0
		for i := 0; i < row; i++ {
			dr := row - 1 - i
			dp[i][j] += up
			dp[dr][j] += down
			if grid[i][j] == 'W' {
				up = 0
			} else if grid[i][j] == 'E' {
				up++
			}
			if grid[dr][j] == 'W' {
				down = 0
			} else if grid[dr][j] == 'E' {
				down++
			}
		}
	}

	res := 0
	for i := range dp {
		for j := range dp[i] {
			if grid[i][j] == '0' {
				res = max(res, dp[i][j])
			}
		}
	}
	return res
}
