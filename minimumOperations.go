package main

// Leetcode LCP.19. (medium)
func minimumOperations(leaves string) int {
	dp := make([][]int, 3)
	for i := range dp {
		dp[i] = make([]int, len(leaves))
	}

	if leaves[0] == 'r' {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}
	dp[1][0] = 100010
	dp[2][0] = 100010

	for i := 1; i < len(leaves); i++ {
		if leaves[i] == 'r' {
			dp[0][i] = dp[0][i-1]
		} else {
			dp[0][i] = dp[0][i-1] + 1
		}
	}
	for i := 1; i < len(leaves); i++ {
		if leaves[i] == 'r' {
			dp[1][i] = min(dp[0][i-1]+1, dp[1][i-1]+1)
		} else {
			dp[1][i] = min(dp[0][i-1], dp[1][i-1])
		}
	}
	for i := 1; i < len(leaves); i++ {
		if leaves[i] == 'r' {
			dp[2][i] = min(dp[1][i-1], dp[2][i-1])
		} else {
			dp[2][i] = min(dp[1][i-1]+1, dp[2][i-1]+1)
		}
	}
	return dp[2][len(leaves)-1]
}
