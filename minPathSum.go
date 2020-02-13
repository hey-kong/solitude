package main

// Leetcode 64. (medium)
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}

	mr, mc := len(grid), len(grid[0])
	for i := 1; i < mr; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < mc; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < mr; i++ {
		for j := 1; j < mc; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[mr-1][mc-1]
}
