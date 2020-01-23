package main

// Leetcode 72. (hard)
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			replace := dp[i-1][j-1] + 1
			if word1[i-1] == word2[j-1] {
				replace = dp[i-1][j-1]
			}
			del := dp[i-1][j] + 1
			insert := dp[i][j-1] + 1
			dp[i][j] = min(replace, min(del, insert))
		}
	}
	return dp[len(word1)][len(word2)]
}
