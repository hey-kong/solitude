package main

// Leetcode m08.11. (medium)
func waysToChange(n int) int {
	coins := []int{1, 5, 10, 25}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < 4; i++ {
		for j := 0; j <= n; j++ {
			if k := j - coins[i]; k >= 0 {
				dp[j] += dp[j-coins[i]]
				dp[j] %= 1000000007
			}
		}
	}
	return dp[n]
}
