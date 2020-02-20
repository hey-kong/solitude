package main

// Leetcode 343. (medium)
func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	}
	if n == 4 {
		return 4
	}
	if n == 5 {
		return 6
	}
	if n == 6 {
		return 9
	}

	dp := make([]int, 3)
	dp[0], dp[1], dp[2] = 4, 6, 9
	for i := 6; i < n; i++ {
		dp[0] = max(3*dp[0], 2*dp[1])
		dp[0], dp[1] = dp[1], dp[0]
		dp[1], dp[2] = dp[2], dp[1]
	}
	return dp[2]
}
