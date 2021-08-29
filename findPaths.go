package main

// Leetcode 576. (medium)
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][][]int, maxMove+1)
	for k := range dp {
		dp[k] = make([][]int, m)
		for i := range dp[k] {
			dp[k][i] = make([]int, n)
		}
	}

	dp[0][startRow][startColumn] = 1
	directions := [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}
	res := 0
	for k := 0; k < maxMove; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if dp[k][i][j] == 0 {
					continue
				}

				for _, d := range directions {
					x, y := i+d[0], j+d[1]
					if x >= 0 && x < m && y >= 0 && y < n {
						dp[k+1][x][y] += dp[k][i][j]
						dp[k+1][x][y] %= 1e9 + 7
					} else {
						res += dp[k][i][j]
						res %= 1e9 + 7
					}
				}
			}
		}
	}
	return res
}
