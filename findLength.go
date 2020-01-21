package main

// Leetcode 718. (medium)
func findLength(A []int, B []int) int {
	al, bl := len(A), len(B)
	dp := make([][]int, al+1)
	for i := range dp {
		dp[i] = make([]int, bl+1)
	}
	maxLength := 0
	for i := 1; i <= al; i++ {
		for j := 1; j <= bl; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				maxLength = max(maxLength, dp[i][j])
			}
		}
	}
	return maxLength
}
