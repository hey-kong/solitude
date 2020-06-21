package main

// Leetcode 10. (hard)
func isMatch(s string, p string) bool {
	s = " " + s
	p = " " + p
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(p))
	}
	dp[0][0] = true
	for i := 1; i < len(p); i++ {
		if p[i] == '*' && dp[0][i-2] {
			dp[0][i] = true
		}
	}

	for i := 1; i < len(s); i++ {
		for j := 1; j < len(p); j++ {
			if p[j] == s[i] || p[j] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j] == '*' {
				if s[i] != p[j-1] && p[j-1] != '.' {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				}
			}
		}
	}
	return dp[len(s)-1][len(p)-1]
}
