package main

// Leetcode 120. (medium)
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}

	res := dp[0]
	for i := 1; i < n; i++ {
		res = min(res, dp[i])
	}
	return res
}
