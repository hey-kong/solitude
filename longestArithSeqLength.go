package main

// Leetcode 1027. (medium)
func longestArithSeqLength(A []int) int {
	dp := make([]map[int]int, len(A))
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	res := 2
	for i := 1; i < len(A); i++ {
		for j := 0; j < i; j++ {
			d := A[i] - A[j]
			if cnt, ok := dp[j][d]; ok {
				dp[i][d] = cnt + 1
				res = max(res, dp[i][d])
			} else {
				dp[i][d] = 2
			}
		}
	}
	return res
}
