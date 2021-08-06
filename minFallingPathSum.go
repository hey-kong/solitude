package main

// Leetcode 931. (medium)
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([]int, n)
	copy(dp, matrix[0])
	for i := 1; i < n; i++ {
		last := dp[0]
		dp[0] = min(dp[0], dp[1]) + matrix[i][0]
		for j := 1; j < n-1; j++ {
			last, dp[j] = dp[j], min(min(last, dp[j]), dp[j+1])+matrix[i][j]
		}
		dp[n-1] = min(last, dp[n-1]) + matrix[i][n-1]
	}

	res := dp[0]
	for i := 1; i < n; i++ {
		res = min(res, dp[i])
	}
	return res
}
