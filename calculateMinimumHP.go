package main

// Leetcode 174. (hard)
func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	dp[n-1][m-1] = 1 - min(0, dungeon[n-1][m-1])
	for i := n - 2; i >= 0; i-- {
		dp[i][m-1] = max(1, dp[i+1][m-1]-dungeon[i][m-1])
	}
	for i := m - 2; i >= 0; i-- {
		dp[n-1][i] = max(1, dp[n-1][i+1]-dungeon[n-1][i])
	}
	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			dp[i][j] = max(1, (min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]))
		}
	}
	return dp[0][0]
}
