package main

// Leetcode 1035. (medium)
func maxUncrossedLines(A []int, B []int) int {
	dp := make([]int, len(B)+1)
	var leftUp int
	for i := range A {
		leftUp = 0
		for j := range B {
			if A[i] == B[j] {
				dp[j+1], leftUp = leftUp+1, dp[j+1]
			} else {
				dp[j+1], leftUp = max(dp[j], dp[j+1]), dp[j+1]
			}
		}
	}
	return dp[len(B)]
}
