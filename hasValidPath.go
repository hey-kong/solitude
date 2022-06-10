package main

// Leetcode 2267. (hard)
func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	if grid[0][0] == ')' || grid[m-1][n-1] == '(' || (m+n-1)%2 != 0 {
		return false
	}

	dp := make([][][]bool, m)
	for i := range dp {
		dp[i] = make([][]bool, n)
		for j := range dp[i] {
			dp[i][j] = make([]bool, (m+n-1)/2+1)
		}
	}

	dp[0][0][1] = true
	for i := range grid {
		for j := range grid[i] {
			delta := 1
			if grid[i][j] == ')' {
				delta = -1
			}
			for k := range dp[i][j] {
				q := k + delta
				if q < 0 || q > (m+n-1)/2 {
					continue
				}
				if i != 0 {
					dp[i][j][q] = dp[i][j][q] || dp[i-1][j][k]
				}
				if j != 0 {
					dp[i][j][q] = dp[i][j][q] || dp[i][j-1][k]
				}
			}
		}
	}
	return dp[m-1][n-1][0]
}
