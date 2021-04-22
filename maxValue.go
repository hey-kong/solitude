package main

// Leetcode m47. (medium)
func maxValue(grid [][]int) int {
	dp := make([][]int, len(grid)+1)
	for i := range dp {
		dp[i] = make([]int, len(grid[0])+1)
	}
	for i := 1; i <= len(grid); i++ {
		for j := 1; j <= len(grid[0]); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[len(grid)][len(grid[0])]
}
