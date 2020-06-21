package main

// Leetcode 629. (hard)
func kInversePairs(n int, k int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		dp[i][0] = 1
		for j := 1; j <= k && j <= i*(i-1)/2; j++ {
			val := dp[i-1][j] + 1000000007
			if j-i >= 0 {
				val -= dp[i-1][j-i]
			}
			val %= 1000000007
			dp[i][j] = (dp[i][j-1] + val) % 1000000007
		}
	}
	return dp[n][k]
}
