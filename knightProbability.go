package main

// Leetcode 688. (medium)
func knightProbability(N int, K int, r int, c int) (res float64) {
	dr := []int{2, 1, -1, -2, -2, -1, 1, 2}
	dc := []int{1, 2, 2, 1, -1, -2, -2, -1}
	dp := make([][]float64, N)
	for i := range dp {
		dp[i] = make([]float64, N)
	}
	dp[r][c] = 1.0

	for step := 0; step < K; step++ {
		nextDp := make([][]float64, N)
		for i := range nextDp {
			nextDp[i] = make([]float64, N)
		}
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				for k := 0; k < 8; k++ {
					tmpR, tmpC := i+dr[k], j+dc[k]
					if tmpR >= 0 && tmpR < N && tmpC >= 0 && tmpC < N {
						nextDp[tmpR][tmpC] += dp[i][j] / 8.0
					}
				}
			}
		}
		dp = nextDp
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			res += dp[i][j]
		}
	}
	return res
}
