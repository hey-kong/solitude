package main

// Leetcode 650. (medium)
func minSteps(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			if i%j == 0 {
				dp[i] = dp[j] + i/j
			}
		}
	}
	return dp[n]
}
