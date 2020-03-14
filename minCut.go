package main

// Leetcode 132. (hard)
func minCut(s string) int {
	isPalindrome := make([][]bool, len(s))
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, len(s))
	}
	for l := 1; l <= len(s); l++ {
		for i := 0; i+l-1 < len(s); i++ {
			if s[i] == s[i+l-1] && (l <= 3 || isPalindrome[i+1][i+l-2]) {
				isPalindrome[i][i+l-1] = true
			}
		}
	}
	dp := make([]int, len(s))
	for i := range s {
		if !isPalindrome[0][i] {
			dp[i] = i
			for j := 0; j < i; j++ {
				if isPalindrome[j+1][i] {
					dp[i] = min(dp[i], dp[j]+1)
				}
			}
		}
	}
	return dp[len(s)-1]
}
