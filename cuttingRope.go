package main

// Leetcode 14- I. (medium)
func cuttingRope(n int) int {
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
	for i := 7; i <= n; i++ {
		dp[0] = max(dp[0]*3, dp[1]*2)
		dp[0], dp[1], dp[2] = dp[1], dp[2], dp[0]
	}
	return dp[2]
}
