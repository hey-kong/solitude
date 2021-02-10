package main

// Leetcode 651. (medium)
func maxA(N int) int {
	if N <= 3 {
		return N
	}

	dp := make([]int, N+1)
	dp[1], dp[2], dp[3] = 1, 2, 3
	for i := 4; i <= N; i++ {
		dp[i] = dp[i-1] + 1
		for j := 1; j <= i-3; j++ {
			dp[i] = max(dp[i], dp[j]*(i-j-1))
		}
	}
	return dp[N]
}
