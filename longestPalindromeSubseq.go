package main

// Leetcode 516. (medium)
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	for l := 2; l <= len(s); l++ {
		for i := 0; i+l-1 < len(s); i++ {
			if s[i] == s[i+l-1] {
				if l == 2 {
					dp[i][i+l-1] = 2
				} else {
					dp[i][i+l-1] = dp[i+1][i+l-2] + 2
				}
			} else {
				dp[i][i+l-1] = max(dp[i+1][i+l-1], dp[i][i+l-2])
			}
		}
	}
	return dp[0][len(s)-1]
}
