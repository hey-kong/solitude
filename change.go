package main

// Leetcode 518. (medium)
func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}

	dp := make([][]int, len(coins))
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := 0; i <= amount; i++ {
		if i-coins[0] >= 0 {
			dp[0][i] = dp[0][i-coins[0]]
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if j-coins[i] >= 0 {
				dp[i][j] += dp[i][j-coins[i]]
			}
		}
	}
	return dp[len(coins)-1][amount]
}
